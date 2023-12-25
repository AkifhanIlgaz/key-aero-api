package controllers

import (
	"net/http"

	"github.com/AkifhanIlgaz/key-aero-api/errors"
	"github.com/AkifhanIlgaz/key-aero-api/models"
	"github.com/AkifhanIlgaz/key-aero-api/services"
	"github.com/AkifhanIlgaz/key-aero-api/utils"
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
	if val, ok := ctx.Get("user"); ok {
		if user, ok := val.(*models.User); ok {
			utils.ResponseWithMessage(ctx, http.StatusOK, gin.H{
				"user": user,
			})
			return
		}
	}

	utils.ResponseWithMessage(ctx, http.StatusUnauthorized, gin.H{
		"message": errors.ErrNotLoggedIn.Error(),
	})
	return
}
