package router

import (
	"Nunu_frame/internal/handler"
	"Nunu_frame/pkg/jwt"
	"Nunu_frame/pkg/log"

	"github.com/spf13/viper"
)

type RouterDeps struct {
	Logger             *log.Logger
	Config             *viper.Viper
	JWT                *jwt.JWT
	UserHandler        *handler.UserHandler
	LabelHandler       *handler.LabelHandler
	CategoryHandler    *handler.CategoryHandler
	SubCategoryHandler *handler.SubCategoryHandler
	MaterialRouter     *handler.MaterialHandler
}
