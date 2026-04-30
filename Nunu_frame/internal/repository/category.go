package repository

import (
	"Nunu_frame/internal/model"
	"context"

	"github.com/gin-gonic/gin"
)

type CategoryRepository interface {
	GetCategory(ctx context.Context, id int64) (*model.Category, error)
	GetAllCategories(ctx *gin.Context) ([]*model.Category, error)
}

func NewCategoryRepository(
	repository *Repository,
) CategoryRepository {
	return &categoryRepository{
		Repository: repository,
	}
}

type categoryRepository struct {
	*Repository
}

func (r *categoryRepository) GetCategory(ctx context.Context, id int64) (*model.Category, error) {
	var category model.Category

	return &category, nil
}

func (r *categoryRepository) GetAllCategories(ctx *gin.Context) ([]*model.Category, error) {
	var categories []*model.Category
	if err := r.DB(ctx).Preload("SubCategories").Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}
