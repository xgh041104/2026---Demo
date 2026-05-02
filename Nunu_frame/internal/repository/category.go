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
	DeleteCategory(ctx *gin.Context, id int64) error
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

// 实现删除接口
func (r *categoryRepository) DeleteCategory(ctx *gin.Context, id int64) error {
	// 查找分类是否存在
	var category model.Category
	if err := r.DB(ctx).Preload("SubCategories").First(&category, id).Error; err != nil {
		return v1.ErrCategoryNotExist
	}

	// 删除该分类下的所有子分类（通过 category_id 字段）
	if err := r.DB(ctx).Where("category_id = ?", id).Delete(&model.SubCategory{}).Error; err != nil {
		return v1.ErrDeleteSubCategory
	}

	// 删除分类（软删除）
	if err := r.DB(ctx).Delete(&category).Error; err != nil {
		return v1.ErrDeleteCategory
	}

	return nil
}
