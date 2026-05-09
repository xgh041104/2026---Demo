package handler

import (
	v1 "Nunu_frame/api/v1"
	"Nunu_frame/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type LabelHandler struct {
	*Handler
	labelService service.LabelService
}

func NewLabelHandler(
	handler *Handler,
	labelService service.LabelService,
) *LabelHandler {
	return &LabelHandler{
		Handler:      handler,
		labelService: labelService,
	}
}

func (h *LabelHandler) GetLabel(ctx *gin.Context) {
	//写出指定查询

}

func (h *LabelHandler) GetAllLabels(ctx *gin.Context) {
	labels, err := h.labelService.GetAllLabels(ctx)
	if err != nil {
		h.logger.Error("get all labels err")
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	v1.HandleSuccess(ctx, labels)
}

func (h *LabelHandler) CreateLabel(ctx *gin.Context) {
	//创建标签
	var req v1.CreateLabelReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		h.logger.Error("create label err bind")
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBind, nil)
		return
	}
	err := h.labelService.CreateLabel(ctx, &req)
	if err != nil {
		h.logger.Error("create label err")
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	v1.HandleSuccess(ctx, nil)
}

// 删除标签
func (h *LabelHandler) DeleteLabel(ctx *gin.Context) {
	LabelId := ctx.Param("id")
	id, err := strconv.ParseInt(LabelId, 10, 64)
	if err != nil {
		h.logger.Error("invalid label id")
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}
	err = h.labelService.DeleteLabel(ctx, id)
	if err != nil {
		h.logger.Error("delete label service error")
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	v1.HandleSuccess(ctx, nil)
}
