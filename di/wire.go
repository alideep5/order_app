//go:build wireinject
// +build wireinject

package di

import (
	"order_app/internal/controller"

	"github.com/google/wire"
)

var AllPublicDeps = wire.NewSet(controller.UserControllerSet)

type AppDi struct {
	UserController controller.UserController
}

func NewAppDi(userController controller.UserController) *AppDi {
	return &AppDi{
		UserController: userController,
	}
}

func InitAppDi() *AppDi {
	wire.Build(NewAppDi, AllPublicDeps)

	return &AppDi{}
}
