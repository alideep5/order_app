package main

import (
	"order_app/di"

	"github.com/gin-gonic/gin"
)

func main() {
	appDi := di.InitAppDi()

	router := gin.Default()

	appDi.UserController.RegisterRoutes(router)

	router.Run(":8080")
}
