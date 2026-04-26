package repository

import (
	"Nunu_frame/internal/model"
	"context"
)

type UserRepository interface {
	FindUserByAccount(c context.Context, req *model.User) (*model.User, bool, error)
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

func (r *userRepository) FindUserByAccount(c context.Context, req *model.User) (*model.User, bool, error) {
	return nil, false, nil
}
