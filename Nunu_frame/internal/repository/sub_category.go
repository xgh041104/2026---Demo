package repository

import (
	"Nunu_frame/internal/model"
	"context"

	"github.com/gin-gonic/gin"
)

type SubCategoryRepository interface {
	GetSubCategory(ctx context.Context, id int64) (*model.SubCategory, error)
	GetAllSubCategories(ctx *gin.Context) ([]*model.SubCategory, error)
}

func NewSubCategoryRepository(
	repository *Repository,
) SubCategoryRepository {
	return &subCategoryRepository{
		Repository: repository,
	}
}

type subCategoryRepository struct {
	*Repository
}

func (r *subCategoryRepository) GetSubCategory(ctx context.Context, id int64) (*model.SubCategory, error) {
	var subCategory model.SubCategory

	return &subCategory, nil
}

func (r *subCategoryRepository) GetAllSubCategories(ctx *gin.Context) ([]*model.SubCategory, error) {
	var subCategories []*model.SubCategory
	if err := r.DB(ctx).Find(&subCategories).Error; err != nil {
		return nil, err
	}
	return subCategories, nil
}
