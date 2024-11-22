package service

import (
	"order_app/internal/persistence/repository"

	"github.com/google/wire"
)

var UserServiceSet = wire.NewSet(NewUserService, repository.UserRepoSet)
