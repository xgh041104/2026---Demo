package service

import (
	v1 "Nunu_frame/api/v1"
	"Nunu_frame/internal/model"
	"Nunu_frame/internal/repository"

	"github.com/gin-gonic/gin"
)

type LabelService interface {
	GetLabel(ctx *gin.Context, id int64) (*model.Label, error)
	GetAllLabels(ctx *gin.Context) ([]*model.Label, error)
	CreateLabel(ctx *gin.Context, req *v1.CreateLabelReq) error
	DeleteLabel(ctx *gin.Context, id int64) error
}

func NewLabelService(
	service *Service,
	labelRepository repository.LabelRepository,
) LabelService {
	return &labelService{
		Service:         service,
		labelRepository: labelRepository,
	}
}

type labelService struct {
	*Service
	labelRepository repository.LabelRepository
}

func (s *labelService) GetLabel(ctx *gin.Context, id int64) (*model.Label, error) {
	return s.labelRepository.GetLabel(ctx, id)
}

func (s *labelService) GetAllLabels(ctx *gin.Context) ([]*model.Label, error) {
	labels, err := s.labelRepository.GetAllLabels(ctx)
	if err != nil {
		return nil, v1.ErrFindLabel
	}
	return labels, nil
}

func (s *labelService) CreateLabel(ctx *gin.Context, req *v1.CreateLabelReq) error {
	//创建标签
	err := s.labelRepository.CreateLabel(ctx, &model.Label{
		Name: req.Name,
	})
	if err != nil {
		return v1.ErrCreateLabel
	}
	return nil
}

// 删除标签
func (s *labelService) DeleteLabel(ctx *gin.Context, id int64) error {
	//删除标签
	err := s.labelRepository.DeleteLabel(ctx, id)
	if err != nil {
		return v1.ErrDeleteLabel
	}
	return nil
}
