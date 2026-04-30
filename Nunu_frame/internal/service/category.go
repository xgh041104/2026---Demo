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

func (s *categoryService) GetAllCategories(ctx *gin.Context) ([]*model.Category, error) {
	categories, err := s.categoryRepository.GetAllCategories(ctx)
	if err != nil {
		return nil, v1.ErrFindCategory
	}
	return categories, nil
}
