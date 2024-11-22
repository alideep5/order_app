package main

import (
	"fmt"
	"order_app/di"
	"order_app/internal/model"
)

func main() {
	appDi := di.InitAppDi()

	newUser := appDi.UserController.CreateAccount(&model.CreateUserDTO{
		Username: "User X",
	})

	fmt.Println("New user user created. Id:", newUser.ID, ", Name:", newUser.Username)

	newUser = appDi.UserController.CreateAccount(&model.CreateUserDTO{
		Username: "USER Y",
	})

	fmt.Println("New user user created. Id:", newUser.ID, ", Name:", newUser.Username)
}
