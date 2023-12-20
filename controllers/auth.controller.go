package controllers

import (
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
	}

	// TODO: Get user

	// TODO: Verify password

	// TODO: Generate tokens

	// TODO: Set cookies

	// TODO: Return with success
}
