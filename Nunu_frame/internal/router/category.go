package router

import "github.com/gin-gonic/gin"

func InitCategoryRouter(
	deps RouterDeps,
	r *gin.RouterGroup,
) {
	noAuthRouter := r.Group("/")
	{
		noAuthRouter.GET("/category", deps.CategoryHandler.GetAllCategories)
	}
}
