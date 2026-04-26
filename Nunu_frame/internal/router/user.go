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
	user := r.Group("/user")
	{
		user.POST("/login")
	}
	userAuth := r.Group("/user").Use(middleware.LoginRequired(deps.JWT, deps.Logger))
	{
		userAuth.GET("/info")
		// Keep for compatibility
	}
	// Strict permission routing group
	admin := r.Group("/admin").Use(middleware.RoleRequired(deps.JWT, deps.Logger, "admin", "root"))
	{
		admin.POST("/createUser")
	}
	root := r.Group("/root").Use(middleware.RoleRequired(deps.JWT, deps.Logger, "root"))
	{
		root.DELETE("/deleteUser/:id")
	}
}
