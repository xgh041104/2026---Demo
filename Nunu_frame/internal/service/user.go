package service

import (
	v1 "Nunu_frame/api/v1"
	"Nunu_frame/internal/repository"
	"context"
)

type UserService interface {
	Login(ctx context.Context, req *v1.LoginRequest) (*v1.LoginResponse, error)
}

func NewUserService(
	service *Service,
	userRepo repository.UserRepository,
) UserService {
	return &userService{
		userRepo: userRepo,
		Service:  service,
	}
}

type userService struct {
	userRepo repository.UserRepository
	*Service
}

func (s *userService) Login(ctx context.Context, req *v1.LoginRequest) (*v1.LoginResponse, error) {
	return nil, nil
}
