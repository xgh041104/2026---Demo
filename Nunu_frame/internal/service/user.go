package service

import (
	v1 "Nunu_frame/api/v1"
	"Nunu_frame/internal/model"
	"Nunu_frame/internal/repository"
	"context"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Login(ctx context.Context, req *v1.LoginRequest) (*v1.LoginResponse, error)
	CreateUser(ctx context.Context, req *v1.CreateUserRequest) error
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
	//判断账号密码是否正确
	if req.Account == "" || req.Password == "" {
		return nil, v1.ErrPasswordNotNil
	}
	//判断用户是否存在
	user, isBool, err := s.userRepo.GetUserByAccount(ctx, req.Account)
	if err != nil {
		return nil, v1.ErrFindUser
	}
	if !isBool {
		return nil, v1.ErrFindUser
	}
	//密码校验
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, v1.ErrPasswordIncorrect
	}

	//生成token
	token, _ := s.jwt.GenToken(user.ID, time.Now().Add(time.Hour*24*7), user.Type)

	//返回token
	return &v1.LoginResponse{
		Token: token,
	}, nil
}

//实现创建用户

func (s *userService) CreateUser(ctx context.Context, req *v1.CreateUserRequest) error {
	//判断用户是否存在
	_, isBool, err := s.userRepo.GetUserByAccount(ctx, req.Account)
	if err != nil {
		return v1.ErrFindUser
	}
	if isBool {
		return v1.ErrUserAlreadyExist
	}
	//密码加密存储
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return v1.ErrPasswordEncrypt
	}

	//创建用户
	return s.userRepo.CreateUser(ctx, &model.User{
		Account:  req.Account,
		Password: string(hashedPassword),
	})
}

// 实现修改接口
