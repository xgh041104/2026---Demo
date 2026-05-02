package handler

import (
	v1 "Nunu_frame/api/v1"
	"Nunu_frame/internal/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SubCategoryHandler struct {
	*Handler
	subCategoryService service.SubCategoryService
}

func NewSubCategoryHandler(
	handler *Handler,
	subCategoryService service.SubCategoryService,
) *SubCategoryHandler {
	return &SubCategoryHandler{
		Handler:            handler,
		subCategoryService: subCategoryService,
	}
}

func (h *SubCategoryHandler) GetSubCategory(ctx *gin.Context) {

}

// 查询接口
func (h *SubCategoryHandler) GetAllSubCategories(ctx *gin.Context) {
	subCategories, err := h.subCategoryService.GetAllSubCategories(ctx)
	if err != nil {
		h.logger.Error("get all sub categories err")
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	v1.HandleSuccess(ctx, subCategories)
}

// 新增接口
func (h *SubCategoryHandler) CreateSubCategory(ctx *gin.Context) {
	var req v1.CreateSubCategoryReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		h.logger.Error("create sub category err bind")
		v1.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}
	err := h.subCategoryService.CreateSubCategory(ctx, &req)
	if err != nil {
		h.logger.Error("create sub category err")
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	v1.HandleSuccess(ctx, nil)
}

// 删除接口
func (h *SubCategoryHandler) DeleteSubCategory(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		h.logger.Error("delete sub category missing id parameter")
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBind, nil)
		return
	}
	var intID int64
	if _, err := fmt.Sscanf(id, "%d", &intID); err != nil {
		h.logger.Error("delete sub category invalid id format")
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBind, nil)
		return
	}

	err := h.subCategoryService.DeleteSubCategory(ctx, intID)
	if err != nil {
		h.logger.Error("delete sub category err")
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	v1.HandleSuccess(ctx, nil)
}
