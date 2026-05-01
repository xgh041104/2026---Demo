package repository

import (
	v1 "Nunu_frame/api/v1"
	"Nunu_frame/internal/model"
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
)

type CategoryRepository interface {
	GetCategory(ctx context.Context, id int64) (*model.Category, error)
	GetAllCategories(ctx *gin.Context) ([]*model.Category, error)
	CreateCategory(ctx *gin.Context, m *model.Category) error
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

// 查询接口
func (r *categoryRepository) GetAllCategories(ctx *gin.Context) ([]*model.Category, error) {
	var categories []*model.Category
	if err := r.DB(ctx).Preload("SubCategories").Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

// 新增接口
func (r *categoryRepository) CreateCategory(ctx *gin.Context, m *model.Category) error {
	var existingCategory model.Category
	if err := r.DB(ctx).Where("name = ?", m.Name).First(&existingCategory).Error; err == nil {
		return v1.ErrCategoryExists
	}

	//检查子分类
	for _, subCategory := range m.SubCategories {
		var existingSubCategory model.SubCategory
		if err := r.DB(ctx).Where("name = ?", subCategory.Name).First(&existingSubCategory).Error; err == nil {
			return fmt.Errorf("子分类名称 '%s' 已存在", subCategory.Name)
		}
	}

	if err := r.DB(ctx).Create(m).Error; err != nil {
		return err
	}

	return nil
}
