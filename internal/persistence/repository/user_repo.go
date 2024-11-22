package repository

import "fmt"

type UserRepo interface {
	CreateUser()
}

type userRepo struct{}

func NewUserRepository() UserRepo {
	return &userRepo{}
}

func (ur *userRepo) CreateUser() {
	fmt.Println("User created")
}
