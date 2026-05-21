//go:build wireinject
// +build wireinject

package wire

import (
	"Nunu_frame/internal/handler"
	"Nunu_frame/internal/repository"
	"Nunu_frame/internal/router"
	"Nunu_frame/internal/server"
	"Nunu_frame/internal/service"
	"Nunu_frame/pkg/app"
	"Nunu_frame/pkg/jwt"
	"Nunu_frame/pkg/log"
	"Nunu_frame/pkg/server/http"

	"github.com/google/wire"
	"github.com/spf13/viper"
)

// ... existing code ...

var repositorySet = wire.NewSet(
	repository.NewDB,
	//repository.NewRedis,
	//repository.NewMongo,
	repository.NewRepository,
	repository.NewTransaction,
	repository.NewUserRepository,
	repository.NewLabelRepository,
	repository.NewCategoryRepository,
	repository.NewSubCategoryRepository,
	repository.NewMaterialRepository,
)

var serviceSet = wire.NewSet(
	service.NewService,
	service.NewUserService,
	service.NewLabelService,
	service.NewCategoryService,
	service.NewSubCategoryService,
	service.NewMaterialService,
)

var handlerSet = wire.NewSet(
	handler.NewHandler,
	handler.NewUserHandler,
	handler.NewLabelHandler,
	handler.NewCategoryHandler,
	handler.NewSubCategoryHandler,
	handler.NewMaterialHandler,
)

// ... existing code ...

var serverSet = wire.NewSet(
	server.NewHTTPServer,
)

// build App
func newApp(
	httpServer *http.Server,
	// task *server.Task,
) *app.App {
	return app.NewApp(
		app.WithServer(httpServer),
		app.WithName("demo-server"),
	)
}

func NewWire(*viper.Viper, *log.Logger) (*app.App, func(), error) {
	panic(wire.Build(
		repositorySet,
		serviceSet,
		handlerSet,
		serverSet,
		wire.Struct(new(router.RouterDeps), "*"),
		jwt.NewJwt,
		newApp,
	))
}
