package router

import (
	"Nunu_frame/internal/middleware"

	"github.com/gin-gonic/gin"
)

func InitUserRouter(
	deps RouterDeps,
	r *gin.RouterGroup,
) {
	// No route group has permission
	noAuthRouter := r.Group("/")
	{
		noAuthRouter.POST("/loginUser", deps.UserHandler.Login)
	}
	// Non-strict permission routing group
	userAuth := r.Group("/user").Use(middleware.LoginRequired(deps.JWT, deps.Logger))
	{
		userAuth.GET("/user")
		userAuth.PUT("/updateUser/:id", deps.UserHandler.UpdateUser)
	}

	// Strict permission routing group
	admin := r.Group("/admin").Use(middleware.TypeRequired(deps.JWT, deps.Logger, "admin", "root"))
	{
		admin.PUT("/user")
		admin.POST("/CreateUser", deps.UserHandler.CreateUser)
		admin.GET("/GetUser", deps.UserHandler.GetUser)
	}

	root := r.Group("/root").Use(middleware.TypeRequired(deps.JWT, deps.Logger, "root"))
	{
		root.PUT("/user")
		root.DELETE("/user/:id", deps.UserHandler.DeleteUser)
	}
}
