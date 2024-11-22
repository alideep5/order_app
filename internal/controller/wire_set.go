package controller

import (
	"order_app/internal/service"

	"github.com/google/wire"
)

var UserControllerSet = wire.NewSet(NewUserController, service.UserServiceSet)
