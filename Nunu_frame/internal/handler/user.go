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

func (h *UserHandler) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	userId, _ := strconv.ParseInt(id, 10, 64)
	err := h.userService.DeleteUser(ctx, uint(userId))
	if err != nil {
		h.logger.Error("delete user service error")
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	v1.HandleSuccess(ctx, nil)
}

func (h *UserHandler) UpdateUser(ctx *gin.Context) {
	var req v1.UpdateUserRequest
	id := ctx.Param("id")
	userId, _ := strconv.ParseInt(id, 10, 64)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		h.logger.Error("update user err bind")
		v1.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}
	err := h.userService.UpdateUser(ctx, &req, uint(userId))
	if err != nil {
		h.logger.Error("update user service error")
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	v1.HandleSuccess(ctx, nil)
}

func (h *UserHandler) GetUser(ctx *gin.Context) {
	var req v1.FindUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		h.logger.Error("get user err bind")
		v1.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}
	user, err := h.userService.GetUser(ctx, &req)
	if err != nil {
		h.logger.Error("get user service error")
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	v1.HandleSuccess(ctx, user)
}
