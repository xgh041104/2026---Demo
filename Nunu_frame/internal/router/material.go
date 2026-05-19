package router

import (
	// "Nunu_frame/internal/middleware"
	"github.com/gin-gonic/gin"
)

func InitMaterialRouter(
	deps RouterDeps,
	r *gin.RouterGroup,
) {
	// No route group has permission
	noAuthRouter := r.Group("/")
	{
		noAuthRouter.GET("/GetAuditingMaterialList", deps.MaterialRouter.GetAuditingMaterialList)
		noAuthRouter.GET("/GetMaterialList", deps.MaterialRouter.GetMaterialList)
		noAuthRouter.POST("/SaveMaterial", deps.MaterialRouter.SaveloadImage)
		noAuthRouter.GET("/HomeMaterialCount", deps.MaterialRouter.HomeMaterialData)
		noAuthRouter.GET("/MaterialUploadTrend", deps.MaterialRouter.MaterialUploadTrend)
		noAuthRouter.PUT("/UpAuditMaterial", deps.MaterialRouter.UpAuditMaterial)
		noAuthRouter.GET("/GetPcCountData", deps.MaterialRouter.GetPcCountData)
		noAuthRouter.GET("/GetPcUrlDataList", deps.MaterialRouter.GetPcUrlDataList)
		noAuthRouter.DELETE("/DeletePcMaterial", deps.MaterialRouter.DeletePcMaterial)
	}
	// Non-strict permission routing group
	// noStrictAuthRouter := r.Group("/").Use(middleware.NoStrictAuth(deps.JWT, deps.Logger))
	// {
	// 	noStrictAuthRouter.GET("/user12")
	// }

	// // Strict permission routing group
	// strictAuthRouter := r.Group("/").Use(middleware.StrictAuth(deps.JWT, deps.Logger))
	// {
	// 	strictAuthRouter.PUT("/user12")
	// }
}
