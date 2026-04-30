package router

import (
	"github.com/gin-gonic/gin"
)

func InitLabelRouter(
	deps RouterDeps,
	r *gin.RouterGroup,
) {
	noAuthRouter := r.Group("/")
	{
		noAuthRouter.GET("/labels", deps.LabelHandler.GetAllLabels)
		noAuthRouter.POST("/label", deps.LabelHandler.CreateLabel)
	}
}
