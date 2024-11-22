package main

import (
	"order_app/di"
)

func main() {
	appDi := di.InitAppDi()

	appDi.UserController.CreateAccount()
}
