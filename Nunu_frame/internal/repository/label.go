package repository

import (
	v1 "Nunu_frame/api/v1"
	"Nunu_frame/internal/model"
	"context"

	"github.com/gin-gonic/gin"
)

type LabelRepository interface {
	GetLabel(ctx context.Context, id int64) (*model.Label, error)
	GetAllLabels(ctx context.Context) ([]*model.Label, error)
	CreateLabel(ctx *gin.Context, m *model.Label) error
	DeleteLabel(ctx *gin.Context, id int64) error
}

func NewLabelRepository(
	repository *Repository,
) LabelRepository {
	return &labelRepository{
		Repository: repository,
	}
}

type labelRepository struct {
	*Repository
}

func (r *labelRepository) GetLabel(ctx context.Context, id int64) (*model.Label, error) {
	var label model.Label

	return &label, nil
}

func (r *labelRepository) GetAllLabels(ctx context.Context) ([]*model.Label, error) {
	var labels []*model.Label
	if err := r.DB(ctx).Model(&model.Label{}).Find(&labels).Error; err != nil {
		return nil, err
	}
	return labels, nil
}

func (r *labelRepository) CreateLabel(ctx *gin.Context, m *model.Label) error {
	// 检查标签名是否已存在
	var count int64
	if err := r.DB(ctx).Model(&model.Label{}).Where("name = ?", m.Name).Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		return v1.ErrLabelAlreadyExist
	}

	// 标签名不存在，可以创建
	return r.DB(ctx).Create(m).Error
}

func (r *labelRepository) DeleteLabel(ctx *gin.Context, id int64) error {
	// 判断 ID 是否为空
	if id <= 0 {
		return v1.ErrBadRequest
	}

	// 检查标签是否存在
	var count int64
	if err := r.DB(ctx).Model(&model.Label{}).Where("id = ?", id).Count(&count).Error; err != nil {
		return err
	}

	if count == 0 {
		return v1.ErrNotFound
	}

	// 执行删除（软删除）
	return r.DB(ctx).Delete(&model.Label{}, id).Error
}
