package repository

import "github.com/google/wire"

var UserRepoSet = wire.NewSet(NewUserRepository)
