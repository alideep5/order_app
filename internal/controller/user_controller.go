package controller

import (
	"net/http"
	"order_app/internal/model"
	"order_app/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type UserController interface {
	RegisterRoutes(r *gin.Engine)
}

type userController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &userController{
		userService: userService,
	}
}

func (uc *userController) RegisterRoutes(r *gin.Engine) {
	r.POST("/user", uc.createAccount)
}

func (uc *userController) createAccount(c *gin.Context) {
	var userDetail model.CreateUserDTO

	if err := c.ShouldBindJSON(&userDetail); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()

	if err := validate.Struct(userDetail); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newUser := uc.userService.CreateAccount(&userDetail)

	c.JSON(http.StatusOK, newUser)
}
