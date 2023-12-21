package controllers

import (
	"github.com/AkifhanIlgaz/key-aero-api/services"
	"github.com/gin-gonic/gin"
)

type AdminController struct {
	userService  *services.UserService
	tokenService *services.TokenService
}

func NewAdminController(userService *services.UserService, tokenService *services.TokenService) *AdminController {
	return &AdminController{
		userService:  userService,
		tokenService: tokenService,
	}
}

func (controller *AdminController) AddUser(ctx *gin.Context) {

}
