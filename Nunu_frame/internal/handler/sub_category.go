package handler

import (
	v1 "Nunu_frame/api/v1"
	"Nunu_frame/internal/service"
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

func (h *SubCategoryHandler) GetAllSubCategories(ctx *gin.Context) {
	subCategories, err := h.subCategoryService.GetAllSubCategories(ctx)
	if err != nil {
		h.logger.Error("get all sub categories err")
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	v1.HandleSuccess(ctx, subCategories)
}
