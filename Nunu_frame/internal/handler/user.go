package handler

import (
	v1 "Nunu_frame/api/v1"
	"Nunu_frame/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	*Handler
	userService service.UserService
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type CreateUserReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func NewUserHandler(handler *Handler, userService service.UserService) *UserHandler {
	return &UserHandler{
		Handler:     handler,
		userService: userService,
	}
}

// 登录接口
func (h *UserHandler) Login(c *gin.Context) {
	var req (v1.LoginRequest)
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Error("login err bind")
		v1.HandleError(c, http.StatusBadRequest, err, nil)
		return
	}
	user, err := h.userService.Login(c, &req)
	if err != nil {
		h.logger.Error("login err")
		v1.HandleError(c, http.StatusInternalServerError, err, nil)
		return
	}
	v1.HandleSuccess(c, user)
}

// 创建接口
func (h *UserHandler) CreateUser(ctx *gin.Context) {
	var req v1.CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		h.logger.Error("create user err bind")
		v1.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}
	err := h.userService.CreateUser(ctx, &req)
	if err != nil {
		h.logger.Error("create user service error")
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	v1.HandleSuccess(ctx, nil)
}

// 修改接口
func (h *UserHandler) UpdateUser(ctx *gin.Context) {
	var req v1.UpdateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		h.logger.Error("update user err bind")
		v1.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}
	err := h.userService.UpdateUser(ctx, &req)
	if err != nil {
		h.logger.Error("update user service error")
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	v1.HandleSuccess(ctx, nil)
}

// 删除接口
func (h *UserHandler) DeleteUser(ctx *gin.Context) {

	UserId := ctx.Param("id")

	id, err := strconv.ParseInt(UserId, 10, 64)
	if err != nil {
		h.logger.Error("invalid user id")
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrNotFindUser, nil)
		return
	}
	err = h.userService.DeleteUser(ctx, id)
	if err != nil {
		h.logger.Error("delete user service error")
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	v1.HandleSuccess(ctx, nil)
}

// 重置用户密码
func (h *UserHandler) ResetUserPassword(ctx *gin.Context) {
	var req v1.ResetUserPasswordRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		h.logger.Error("update user err bind")
		v1.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}
	err := h.userService.ResetUserPassword(ctx, &req)
	if err != nil {
		h.logger.Error("update user service error")
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	v1.HandleSuccess(ctx, nil)
}
