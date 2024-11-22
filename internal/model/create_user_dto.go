package model

type CreateUserDTO struct {
	Username string `json:"username" validate:"required,min=3,max=20"`
}
