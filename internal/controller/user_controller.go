package controller

import (
	"order_app/internal/model"
	"order_app/internal/service"
)

type UserController interface {
	CreateAccount(createUserDTO *model.CreateUserDTO) *model.UserDetail
}

type userController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &userController{
		userService: userService,
	}
}

func (uc *userController) CreateAccount(createUserDTO *model.CreateUserDTO) *model.UserDetail {
	return uc.userService.CreateAccount(createUserDTO)
}
