package repository

import (
	"Nunu_frame/internal/model"
	"context"
	"errors"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindUserByAccount(ctx context.Context, req *model.User) (*model.User, bool, error)
	GetUserByAccount(ctx context.Context, account string) (*model.User, bool, error)
	CreateUser(ctx context.Context, user *model.User) error
	GetUserById(ctx context.Context, id uint) (*model.User, bool, error)
	DeleteUser(ctx context.Context, id uint) error
	UpdateUser(ctx context.Context, updates map[string]interface{}, id uint) error
	FindUser(ctx context.Context, req *model.User, page, pageSize int) (int64, []*model.User, error)
}

func NewUserRepository(
	r *Repository,
) UserRepository {
	return &userRepository{
		Repository: r,
	}
}

type userRepository struct {
	*Repository
}

func (r *userRepository) FindUserByAccount(ctx context.Context, req *model.User) (*model.User, bool, error) {
	var user model.User
	err := r.DB(ctx).Where("account = ?", req.Account).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, false, nil
		}
		return nil, false, err
	}
	return &user, true, nil
}

// 查询账户名称
func (r *userRepository) GetUserByAccount(ctx context.Context, account string) (*model.User, bool, error) {
	var user model.User
	err := r.DB(ctx).Where("account = ?", account).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, false, nil
		}
		return nil, false, err
	}

	return &user, true, nil
}

// 查询用户ID
func (r *userRepository) GetUserById(ctx context.Context, id uint) (*model.User, bool, error) {
	var user model.User
	err := r.DB(ctx).Where("id = ?", id).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, false, nil
		}
		return nil, false, err
	}

	return &user, true, nil
}

// 创建用户
func (r *userRepository) CreateUser(ctx context.Context, user *model.User) error {
	return r.DB(ctx).Create(user).Error

}

func (r *userRepository) DeleteUser(ctx context.Context, id uint) error {
	return r.DB(ctx).Delete(&model.User{}, id).Error
}

func (r *userRepository) UpdateUser(ctx context.Context, updates map[string]interface{}, id uint) error {
	return r.DB(ctx).Model(&model.User{}).Where("id = ?", id).Updates(updates).Error
}

func (r *userRepository) FindUser(ctx context.Context, req *model.User, page, pageSize int) (int64, []*model.User, error) {
	query := r.DB(ctx).Model(&model.User{})
	if req.Account != "" {
		query = query.Where("account LIKE ?", "%"+req.Account+"%")
	}
	if req.RealName != "" {
		query = query.Where("real_name LIKE ?", "%"+req.RealName+"%")
	}
	if req.Phone != "" {
		query = query.Where("phone LIKE ?", "%"+req.Phone+"%")
	}

	var total int64
	err := query.Count(&total).Error
	if err != nil {
		return 0, nil, err
	}

	var users []*model.User
	err = query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&users).Error
	if err != nil {
		return 0, nil, err
	}
	return total, users, nil
}
