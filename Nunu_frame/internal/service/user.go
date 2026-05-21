package service

import (
	v1 "Nunu_frame/api/v1"
	"Nunu_frame/internal/model"
	"Nunu_frame/internal/repository"
	"context"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Login(ctx context.Context, req *v1.LoginRequest) (*v1.LoginResponse, error)
	CreateUser(ctx context.Context, req *v1.CreateUserRequest) error
	DeleteUser(ctx context.Context, id uint) error
	UpdateUser(ctx context.Context, req *v1.UpdateUserRequest, id uint) error
	GetUser(ctx context.Context, req *v1.FindUserRequest) (*v1.UserListResponse, error)
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

func (s *userService) DeleteUser(ctx context.Context, id uint) error {
	err := s.tm.Transaction(ctx, func(ctx context.Context) error {
		user, isBool, err := s.userRepo.GetUserById(ctx, id)
		if err != nil {
			return v1.ErrFindUser
		}
		if !isBool {
			return v1.ErrFindUser
		}
		if user.IsDisabled == 0 {
			return v1.ErrUserNotDisabled
		}
		return s.userRepo.DeleteUser(ctx, id)
	})
	if err != nil {
		return v1.ErrDeleteUser
	}
	return nil
}

func (s *userService) UpdateUser(ctx context.Context, req *v1.UpdateUserRequest, id uint) error {
	updates := make(map[string]interface{})
	if req.Account != "" {
		updates["account"] = req.Account
	}
	if req.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			return v1.ErrPasswordEncrypt
		}
		updates["password"] = string(hashedPassword)
	}
	if req.RealName != "" {
		updates["real_name"] = req.RealName
	}
	if req.Phone != "" {
		updates["phone"] = req.Phone
	}
	if req.Email != "" {
		updates["email"] = req.Email
	}
	if req.Type != "" {
		updates["type"] = req.Type
	}
	updates["is_disabled"] = req.IsDisabled

	err := s.userRepo.UpdateUser(ctx, updates, id)
	if err != nil {
		return v1.ErrUpdateUser
	}
	return nil
}

func (s *userService) GetUser(ctx context.Context, req *v1.FindUserRequest) (*v1.UserListResponse, error) {

	total, list, err := s.userRepo.FindUser(ctx, &model.User{
		Account:  req.Account,
		RealName: req.RealName,
		Phone:    req.Phone,
	}, req.Page, req.PageSize)
	if err != nil {
		return nil, v1.ErrUserFailFind
	}
	fmt.Println("-------------------------------", total, list)
	responseList := make([]v1.FindUserResponse, 0, len(list))
	for _, user := range list {

		responseList = append(responseList, v1.FindUserResponse{
			Account:  user.Account,
			RealName: user.RealName,
			Phone:    user.Phone,
		})
	}
	return &v1.UserListResponse{
		Total: total,
		List:  responseList,
	}, nil
}
