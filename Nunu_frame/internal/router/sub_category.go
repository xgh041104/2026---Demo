package router

import "github.com/gin-gonic/gin"

func InitSubCategoryRouter(
	deps RouterDeps,
	r *gin.RouterGroup,
) {
	noAuthRouter := r.Group("/")
	{
		noAuthRouter.GET("/sub-categories", deps.SubCategoryHandler.GetAllSubCategories)
	}
}
