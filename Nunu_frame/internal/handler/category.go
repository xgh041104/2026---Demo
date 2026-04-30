package handler

import (
	v1 "Nunu_frame/api/v1"
	"Nunu_frame/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	*Handler
	categoryService service.CategoryService
}

func NewCategoryHandler(
	handler *Handler,
	categoryService service.CategoryService,
) *CategoryHandler {
	return &CategoryHandler{
		Handler:         handler,
		categoryService: categoryService,
	}
}

func (h *CategoryHandler) GetCategory(ctx *gin.Context) {

}

func (h *CategoryHandler) GetAllCategories(ctx *gin.Context) {
	categories, err := h.categoryService.GetAllCategories(ctx)
	if err != nil {
		h.logger.Error("get all categories err")
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	v1.HandleSuccess(ctx, categories)
}
