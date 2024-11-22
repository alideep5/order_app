package repository

import (
	"order_app/internal/model"
	"order_app/internal/persistence/entity"
)

type UserRepo interface {
	CreateUser(createUserDTO *model.CreateUserDTO) *model.UserDetail
}

type userRepo struct {
	users []entity.User
}

func NewUserRepository() UserRepo {
	return &userRepo{}
}

func (ur *userRepo) CreateUser(createUserDTO *model.CreateUserDTO) *model.UserDetail {
	user := entity.User{
		ID:       len(ur.users) + 1,
		Username: createUserDTO.Username,
	}

	ur.users = append(ur.users, user)

	return &model.UserDetail{
		ID:       user.ID,
		Username: user.Username,
	}
}
