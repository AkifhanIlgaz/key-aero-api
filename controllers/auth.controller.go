package controllers

import (
	"fmt"
	"net/http"

	"github.com/AkifhanIlgaz/key-aero-api/cfg"
	"github.com/AkifhanIlgaz/key-aero-api/models"
	"github.com/AkifhanIlgaz/key-aero-api/services"
	"github.com/AkifhanIlgaz/key-aero-api/utils"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	userService  *services.UserService
	tokenService *services.TokenService
	config       *cfg.Config
}

func NewAuthController(config *cfg.Config, userService *services.UserService, tokenService *services.TokenService) *AuthController {
	return &AuthController{
		userService:  userService,
		tokenService: tokenService,
		config:       config,
	}
}

func (controller *AuthController) SignIn(ctx *gin.Context) {
	var credentials models.Credentials
	if err := ctx.ShouldBindJSON(&credentials); err != nil {
		utils.ResponseWithMessage(ctx, http.StatusBadRequest, gin.H{
			"message": "There are some missing required fields",
		})
		return
	}

	// TODO: Get user
	user, err := controller.userService.GetUser(credentials.Username)
	if err != nil {
		utils.ResponseWithMessage(ctx, http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("%v doesnt't exist", credentials.Username),
		})
		return
	}

	// TODO: Verify password
	err = utils.VerifyPassword(user.PasswordHash, credentials.Password)
	if err != nil {
		utils.ResponseWithMessage(ctx, http.StatusUnauthorized, gin.H{
			"message": "Wrong password",
		})
		return
	}

	// TODO: Generate tokens
	accessToken, err := controller.tokenService.GenerateAccessToken(user.Id)
	if err != nil {
		utils.ResponseWithMessage(ctx, http.StatusInternalServerError, gin.H{
			"message": "Something went wrong",
		})
		return
	}

	refreshToken, err := controller.tokenService.GenerateRefreshToken(user.Id)
	if err != nil {
		utils.ResponseWithMessage(ctx, http.StatusInternalServerError, gin.H{
			"message": "Something went wrong",
		})
		return
	}

	// TODO: Return with success
	utils.ResponseWithMessage(ctx, http.StatusOK, gin.H{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	})
}
