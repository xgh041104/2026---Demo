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
func (r *userRepository) GetUserById(ctx context.Context, id int64) (*model.User, bool, error) {
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
