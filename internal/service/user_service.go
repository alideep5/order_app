package service

import "order_app/internal/persistence/repository"

type UserService interface {
	CreateAccount()
}

type userService struct {
	userRepo repository.UserRepo
}

func NewUserService(userRepo repository.UserRepo) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (us *userService) CreateAccount() {
	us.userRepo.CreateUser()
}
