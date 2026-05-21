package handler

import (
	v1 "Nunu_frame/api/v1"
	"Nunu_frame/internal/model"
	"Nunu_frame/internal/service"
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type MaterialHandler struct {
	*Handler
	materialService service.MaterialService
}

func NewMaterialHandler(
	handler *Handler,
	materialService service.MaterialService,
) *MaterialHandler {
	return &MaterialHandler{
		Handler:         handler,
		materialService: materialService,
	}
}

func (h *MaterialHandler) GetAuditingMaterialList(c *gin.Context) {
	var req v1.PageNumAndPageSizeRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		h.logger.Error("GetAuditingMaterialHandler bind json error", zap.Any("error", err))
		v1.HandleError(c, http.StatusBadRequest, v1.ErrBind, nil)
		return
	}
	GAT, err := h.materialService.GetAuditingMaterialList(c, &req)
	if err != nil {
		h.logger.Error("GetAuditingMaterialHandler error", zap.Any("error", err))
		v1.HandleError(c, http.StatusInternalServerError, err, nil)
		return
	}
	v1.HandleSuccess(c, GAT)
}

// 存储素材到本地z
func (h *MaterialHandler) SaveloadImage(c *gin.Context) {
	var path string
	var FileName string
	var url string
	var avatar string
	var (
		imageExts = map[string]bool{
			".jpg":  true,
			".jpeg": true,
			".png":  true,
			".gif":  true,
			".bmp":  true,
			".webp": true,
			".svg":  true,
		}
		videoExts = map[string]bool{
			".mp4":  true,
			".avi":  true,
			".mov":  true,
			".wmv":  true,
			".flv":  true,
			".mkv":  true,
			".webm": true,
		}
	)

	file, err := c.FormFile("file")

	if err == http.ErrMissingFile {
		path = ""
		FileName = ""
		url = ""
		avatar = ""
	} else if err != nil {
		v1.HandleError(c, http.StatusBadRequest, err, nil)
		return
	} else {
		// 获取原始全名
		originName := file.Filename
		// 拆分后缀
		ext := strings.ToLower(filepath.Ext(originName))
		// 去掉后缀的纯原名
		onlyName := strings.TrimSuffix(originName, filepath.Ext(originName))

		if imageExts[ext] {
			avatar = "image"
		} else if videoExts[ext] {
			avatar = "video"
		} else {
			v1.HandleError(c, http.StatusServiceUnavailable, v1.ErrFileExt, nil)
			return
		}

		// 拼接新名：时间戳_原文件名.后缀
		newName := fmt.Sprintf("%d_%s%s", time.Now().Unix(), onlyName, ext)
		FileName = newName
		path = fmt.Sprintf("static/%s/%s", avatar, newName)

		if err := c.SaveUploadedFile(file, path); err != nil {
			v1.HandleError(c, http.StatusServiceUnavailable, err, nil)
			return
		}
		url = "/static/" + avatar + "/" + newName
	}

	OldStatus, _ := strconv.Atoi(c.PostForm("status"))
	title := c.PostForm("title")
	creator_id, _ := strconv.Atoi(c.PostForm("creator_id"))
	label_id := c.PostForm("label_id")
	category_id, _ := strconv.Atoi(c.PostForm("category_id"))
	remark := c.PostForm("remark")

	// 封装
	req := &model.Material{
		Name:       FileName,
		CreatorId:  creator_id,
		ImageUrl:   url,
		LabelId:    label_id,
		CategoryId: category_id,
		Status:     OldStatus,
		Remark:     remark,
	}

	if title != "" {
		req.Name = title
	}

	if err := h.materialService.SaveAvatar(c, req); err != nil {
		v1.HandleError(c, http.StatusBadGateway, err, nil)
		return
	}

	v1.HandleSuccess(c, gin.H{"msg": "success"})
}

// 素材页面 - 搜索
func (h *MaterialHandler) GetMaterialList(c *gin.Context) {
	var req v1.ReqSearchMaterial
	if err := c.ShouldBindQuery(&req); err != nil {
		h.logger.Error("SearchMaterial bind Query error", zap.Any("error", err))
		v1.HandleError(c, http.StatusBadRequest, v1.ErrBind, nil)
		return
	}
	GET, err := h.materialService.GetMaterialList(c, &req)
	if err != nil {
		h.logger.Error("获取素材列表失败", zap.Any("error", err))
		v1.HandleError(c, http.StatusInternalServerError, err, nil)
		return
	}
	v1.HandleSuccess(c, GET)
}

// 首页-素材总数
func (h *MaterialHandler) HomeMaterialData(c *gin.Context) {
	var req v1.PageNumSizeAndStatus
	if err := c.ShouldBindQuery(&req); err != nil {
		h.logger.Error("GetMaterialDataHandler bind Query error", zap.Any("error", err))
		v1.HandleError(c, http.StatusBadRequest, v1.ErrBind, nil)
		return
	}
	Count, err := h.materialService.HomeMaterialData(c, req.Status)
	if err != nil {
		h.logger.Error("获取素材总数失败", zap.Any("error", err))
		v1.HandleError(c, http.StatusInternalServerError, err, nil)
		return
	}
	v1.HandleSuccess(c, Count)
}

// 首页-近7日素材上传趋势
func (h *MaterialHandler) MaterialUploadTrend(c *gin.Context) {
	List, err := h.materialService.MaterialUploadTrend(c)
	if err != nil {
		h.logger.Error("获取素材上传趋势失败", zap.Any("error", err))
		v1.HandleError(c, http.StatusInternalServerError, err, nil)
		return
	}
	v1.HandleSuccess(c, List)
}

func (h *MaterialHandler) UpAuditMaterial(c *gin.Context) {
	var req v1.UpAuditMaterialReq
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Error("GetMaterialUpData bind JSON error", zap.Any("error", err))
		v1.HandleError(c, http.StatusBadRequest, v1.ErrBind, nil)
		return
	}
	err := h.materialService.UpAuditMaterial(c, &req)
	if err != nil {
		h.logger.Error("更新数据失败", zap.Any("error", err))
		v1.HandleError(c, http.StatusInternalServerError, err, nil)
		return
	}
	v1.HandleSuccess(c, gin.H{"status": "更新成功", "id": req.Id})
}

// 个人中心 - 获取个人数据数量
func (h *MaterialHandler) GetPcCountData(c *gin.Context) {
	var req v1.PcData
	if err := c.ShouldBindQuery(&req); err != nil {
		h.logger.Error("GetPcCountData bind Query error", zap.Any("error", err))
		v1.HandleError(c, http.StatusBadRequest, v1.ErrBind, nil)
		return
	}
	Count, err := h.materialService.GetPcCountData(c, &req)
	if err != nil {
		h.logger.Error("获取个人数据失败", zap.Any("error", err))
		v1.HandleError(c, http.StatusInternalServerError, err, nil)
		return
	}

	v1.HandleSuccess(c, gin.H{"count": Count})
}

func (h *MaterialHandler) GetPcUrlDataList(c *gin.Context) {
	var req v1.PcData
	if err := c.ShouldBindQuery(&req); err != nil {
		h.logger.Error("GetPcUrlDataList bind Query error", zap.Any("error", err))
		v1.HandleError(c, http.StatusBadRequest, v1.ErrBind, nil)
		return
	}
	List, err := h.materialService.GetPcUrlDataList(c, &req)
	if err != nil {
		h.logger.Error("获取个人素材列表失败", zap.Any("error", err))
		v1.HandleError(c, http.StatusInternalServerError, err, nil)
		return
	}

	v1.HandleSuccess(c, List)
}

func (h *MaterialHandler) DeletePcMaterial(c *gin.Context) {
	var req v1.PcData
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Error("DeletePcMaterial bind JSON error", zap.Any("error", err))
		v1.HandleError(c, http.StatusBadRequest, v1.ErrBind, nil)
		return
	}
	err := h.materialService.DeletePcMaterial(c, &req)
	if err != nil {
		h.logger.Error("删除个人素材失败", zap.Any("error", err))
		v1.HandleError(c, http.StatusInternalServerError, err, nil)
		return
	}
	v1.HandleSuccess(c, gin.H{"MaterialId": req.Id, "msg": "success"})
}

// 素材页面 - 搜索
// func (h *MaterialHandler) SearchMaterial(c *gin.Context) {
// 	var req v1.ReqSearchMaterial
// 	if err := c.ShouldBindQuery(&req); err != nil {
// 		h.logger.Error("SearchMaterial bind Query error", zap.Any("error", err))
// 		v1.HandleError(c, http.StatusBadRequest, v1.ErrBind, nil)
// 		return
// 	}
// 	List, err := h.materialService.SearchMaterial(c, req)
// 	if err != nil {
// 		h.logger.Error("获取素材列表失败", zap.Any("error", err))
// 		v1.HandleError(c, http.StatusInternalServerError, err, nil)
// 		return
// 	}
// }
