package server

import (
	apiV1 "Nunu_frame/api/v1"
	"Nunu_frame/internal/middleware"
	"Nunu_frame/internal/router"
	"Nunu_frame/pkg/server/http"

	"github.com/gin-gonic/gin"
)

func NewHTTPServer(
	deps router.RouterDeps,
) *http.Server {
	if deps.Config.GetString("env") == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}
	s := http.NewServer(
		gin.Default(),
		deps.Logger,
		http.WithServerHost(deps.Config.GetString("http.host")),
		http.WithServerPort(deps.Config.GetInt("http.port")),
	)

	s.Use(
		middleware.CORSMiddleware(),
		middleware.ResponseLogMiddleware(deps.Logger),
		middleware.RequestLogMiddleware(deps.Logger),
		//middleware.SignMiddleware(log),
	)
	s.GET("/", func(ctx *gin.Context) {
		deps.Logger.WithContext(ctx).Info("hello")
		apiV1.HandleSuccess(ctx, map[string]interface{}{
			":)": "Thank you for using nunu!",
		})
	})

	s.Static("/static", "./static")

	v1 := s.Group("/v1")
	router.InitUserRouter(deps, v1)
	router.InitLabelRouter(deps, v1)
	router.InitCategoryRouter(deps, v1)
	router.InitSubCategoryRouter(deps, v1)
	router.InitMaterialRouter(deps, v1)

	return s
}
