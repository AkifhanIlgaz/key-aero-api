package controllers

import (
	"github.com/AkifhanIlgaz/key-aero-api/services"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (controller *UserController) Me(ctx *gin.Context) {

}
