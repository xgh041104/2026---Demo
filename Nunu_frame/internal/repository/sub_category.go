package repository

import (
	v1 "Nunu_frame/api/v1"
	"Nunu_frame/internal/model"
	"context"

	"github.com/gin-gonic/gin"
)

type SubCategoryRepository interface {
	GetSubCategory(ctx context.Context, id int64) (*model.SubCategory, error)
	GetAllSubCategories(ctx *gin.Context) ([]*model.SubCategory, error)
	CreateSubCategory(ctx *gin.Context, m *model.SubCategory) error
	DeleteSubCategory(ctx *gin.Context, id int64) error
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

// 查询
func (r *subCategoryRepository) GetAllSubCategories(ctx *gin.Context) ([]*model.SubCategory, error) {
	var subCategories []*model.SubCategory
	if err := r.DB(ctx).Find(&subCategories).Error; err != nil {
		return nil, err
	}
	return subCategories, nil
}

// 新增
func (r *subCategoryRepository) CreateSubCategory(ctx *gin.Context, m *model.SubCategory) error {
	// 检查子分类名称是否已存在
	var existingSubCategory model.SubCategory
	if err := r.DB(ctx).Where("name = ?", m.Name).First(&existingSubCategory).Error; err == nil {
		return v1.ErrSubCategoryAlreadyExist
	}

	// 创建子分类
	if err := r.DB(ctx).Create(m).Error; err != nil {
		return err
	}

	return nil
}

// ... existing code ...

// 删除
func (r *subCategoryRepository) DeleteSubCategory(ctx *gin.Context, id int64) error {
	// 检查子分类是否存在
	var subCategory model.SubCategory
	if err := r.DB(ctx).First(&subCategory, id).Error; err != nil {
		return v1.ErrSubCategoryNotExist
	}

	// 删除子分类（软删除）
	if err := r.DB(ctx).Delete(&subCategory).Error; err != nil {
		return v1.ErrDeleteSubCategory
	}

	return nil
}
