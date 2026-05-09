package service

import (
	v1 "Nunu_frame/api/v1"
	"Nunu_frame/internal/model"
	"Nunu_frame/internal/repository"
	"context"

	"github.com/gin-gonic/gin"
)

type SubCategoryService interface {
	GetSubCategory(ctx context.Context, id int64) (*model.SubCategory, error)
	GetAllSubCategories(ctx *gin.Context) ([]*model.SubCategory, error)
	CreateSubCategory(ctx *gin.Context, v *v1.CreateSubCategoryReq) error
	DeleteSubCategory(ctx *gin.Context, id int64) error
}

func NewSubCategoryService(
	service *Service,
	subCategoryRepository repository.SubCategoryRepository,
) SubCategoryService {
	return &subCategoryService{
		Service:               service,
		subCategoryRepository: subCategoryRepository,
	}
}

type subCategoryService struct {
	*Service
	subCategoryRepository repository.SubCategoryRepository
}

func (s *subCategoryService) GetSubCategory(ctx context.Context, id int64) (*model.SubCategory, error) {
	return s.subCategoryRepository.GetSubCategory(ctx, id)
}

// 查询
func (s *subCategoryService) GetAllSubCategories(ctx *gin.Context) ([]*model.SubCategory, error) {
	subCategories, err := s.subCategoryRepository.GetAllSubCategories(ctx)
	if err != nil {
		return nil, v1.ErrFindSubCategory
	}
	return subCategories, nil
}

// 新增
func (s *subCategoryService) CreateSubCategory(ctx *gin.Context, v *v1.CreateSubCategoryReq) error {
	return s.subCategoryRepository.CreateSubCategory(ctx, &model.SubCategory{
		Name:       v.Name,
		Remark:     v.Remark,
		Sort:       v.Sort,
		Status:     v.Status,
		CategoryId: v.CategoryId,
	})
}

// 删除
func (s *subCategoryService) DeleteSubCategory(ctx *gin.Context, id int64) error {
	return s.subCategoryRepository.DeleteSubCategory(ctx, id)
}
