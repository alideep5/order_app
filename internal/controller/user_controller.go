package controller

import "order_app/internal/service"

type UserController interface {
	CreateAccount()
}

type userController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &userController{
		userService: userService,
	}
}

func (uc *userController) CreateAccount() {
	uc.userService.CreateAccount()
}
