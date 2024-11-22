package service

import (
	"order_app/internal/model"
	"order_app/internal/persistence/repository"
)

type UserService interface {
	CreateAccount(createUserDTO *model.CreateUserDTO) *model.UserDetail
}

type userService struct {
	userRepo repository.UserRepo
}

func NewUserService(userRepo repository.UserRepo) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (us *userService) CreateAccount(createUserDTO *model.CreateUserDTO) *model.UserDetail {
	return us.userRepo.CreateUser(createUserDTO)
}
