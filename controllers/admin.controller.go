package controllers

import (
	"net/http"

	"github.com/AkifhanIlgaz/key-aero-api/services"
	"github.com/AkifhanIlgaz/key-aero-api/utils"
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

func (controller *AdminController) GetAllUsers(ctx *gin.Context) {
	users, err := controller.userService.GetUsers()
	if err != nil {
		utils.ResponseWithMessage(ctx, http.StatusInternalServerError, nil)
		return
	}

	utils.ResponseWithMessage(ctx, http.StatusFound, gin.H{
		"users": users,
	})
}

func (controller *AdminController) UpdateUser(ctx *gin.Context) {

}
func (controller *AdminController) DeleteUser(ctx *gin.Context) {

}
