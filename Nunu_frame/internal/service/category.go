package service

import (
	v1 "Nunu_frame/api/v1"
	"Nunu_frame/internal/model"
	"Nunu_frame/internal/repository"
	"context"

	"github.com/gin-gonic/gin"
)

type CategoryService interface {
	GetCategory(ctx context.Context, id int64) (*model.Category, error)
	GetAllCategories(ctx *gin.Context) ([]*model.Category, error)
	CreateCategory(ctx *gin.Context, c *v1.CreateCategoryReq) error
	DeleteCategory(ctx *gin.Context, ID int64) error
}

func NewCategoryService(
	service *Service,
	categoryRepository repository.CategoryRepository,
) CategoryService {
	return &categoryService{
		Service:            service,
		categoryRepository: categoryRepository,
	}
}

type categoryService struct {
	*Service
	categoryRepository repository.CategoryRepository
}

func (s *categoryService) GetCategory(ctx context.Context, id int64) (*model.Category, error) {
	return s.categoryRepository.GetCategory(ctx, id)
}

// 查询全部
func (s *categoryService) GetAllCategories(ctx *gin.Context) ([]*model.Category, error) {
	categories, err := s.categoryRepository.GetAllCategories(ctx)
	if err != nil {
		return nil, v1.ErrFindCategory
	}
	return categories, nil
}

// 新增接口
func (s *categoryService) CreateCategory(ctx *gin.Context, c *v1.CreateCategoryReq) error {
	return s.categoryRepository.CreateCategory(ctx, &model.Category{
		Name:   c.Name,
		Remark: c.Remark,
		Sort:   c.Sort,
		Status: c.Status,
	})
}

// 删除接口
func (s *categoryService) DeleteCategory(ctx *gin.Context, ID int64) error {
	return s.categoryRepository.DeleteCategory(ctx, ID)
}
