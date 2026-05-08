package handler

import (
	v1 "Nunu_frame/api/v1"
	"Nunu_frame/internal/service"
	"fmt"
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

// 新增分类
func (h *CategoryHandler) CreateCategory(ctx *gin.Context) {
	var req v1.CreateCategoryReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		h.logger.Error("create category err bind")
		v1.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}
	err := h.categoryService.CreateCategory(ctx, &req)
	if err != nil {
		h.logger.Error("create category err")
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	v1.HandleSuccess(ctx, nil)
}

// 删除分类
func (h *CategoryHandler) DeleteCategory(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		h.logger.Error("delete category missing id parameter")
		v1.HandleError(ctx, http.StatusBadRequest, nil, nil)
		return
	}

	var intID int64
	if _, err := fmt.Sscanf(id, "%d", &intID); err != nil {
		h.logger.Error("delete category invalid id format")
		v1.HandleError(ctx, http.StatusBadRequest, nil, nil)
		return
	}

	err := h.categoryService.DeleteCategory(ctx, intID)
	if err != nil {
		h.logger.Error("delete category err")
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	v1.HandleSuccess(ctx, nil)
}
