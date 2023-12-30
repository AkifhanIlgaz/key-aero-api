package controllers

import (
	"fmt"
	"net/http"

	"github.com/AkifhanIlgaz/key-aero-api/cfg"
	"github.com/AkifhanIlgaz/key-aero-api/errors"
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

	fmt.Println(credentials)

	user, err := controller.userService.GetUserByUsername(credentials.Username)
	if err != nil {
		fmt.Println(err.Error())
		utils.ResponseWithMessage(ctx, http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("%v doesnt't exist", credentials.Username),
		})
		return
	}

	err = utils.VerifyPassword(user.PasswordHash, credentials.Password)
	if err != nil {
		utils.ResponseWithMessage(ctx, http.StatusUnauthorized, gin.H{
			"message": "Wrong password",
		})
		return
	}

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

	utils.ResponseWithMessage(ctx, http.StatusOK, gin.H{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	})
}

func (controller *AuthController) SignOut(ctx *gin.Context) {
	var credentials models.SignOutCredentials
	if err := ctx.ShouldBindJSON(&credentials); err != nil {
		utils.ResponseWithMessage(ctx, http.StatusBadRequest, gin.H{
			"message": errors.ErrRefreshTokenMissing.Error(),
		})
		return
	}

	_, err := controller.tokenService.DeleteRefreshToken(credentials.RefreshToken)
	if err != nil {
		if err != errors.ErrRefreshTokenMissing {
			utils.ResponseWithMessage(ctx, http.StatusBadRequest, gin.H{
				"message": errors.ErrRefreshTokenMissing.Error(),
			})
			return
		}
		utils.ResponseWithMessage(ctx, http.StatusInternalServerError, gin.H{
			"message": errors.ErrSomethingWentWrong.Error(),
		})
		return
	}

	utils.ResponseWithMessage(ctx, http.StatusOK, gin.H{
		"message": "Logged out successfully",
	})
}

func (controller *AuthController) Refresh(ctx *gin.Context) {
	var credentials models.SignOutCredentials
	if err := ctx.ShouldBindJSON(&credentials); err != nil {
		utils.ResponseWithMessage(ctx, http.StatusBadRequest, gin.H{
			"message": errors.ErrRefreshTokenMissing.Error(),
		})
		return
	}

	uid, err := controller.tokenService.DeleteRefreshToken(credentials.RefreshToken)
	if err != nil {
		if err == errors.ErrRefreshTokenMissing {
			utils.ResponseWithMessage(ctx, http.StatusBadRequest, gin.H{
				"message": errors.ErrRefreshTokenMissing.Error(),
			})
			return
		}
		utils.ResponseWithMessage(ctx, http.StatusInternalServerError, gin.H{
			"message": errors.ErrSomethingWentWrong.Error(),
		})
		return
	}

	accessToken, err := controller.tokenService.GenerateAccessToken(uid)
	if err != nil {
		utils.ResponseWithMessage(ctx, http.StatusInternalServerError, gin.H{
			"message": "Something went wrong",
		})
		return
	}

	refreshToken, err := controller.tokenService.GenerateRefreshToken(uid)
	if err != nil {
		utils.ResponseWithMessage(ctx, http.StatusInternalServerError, gin.H{
			"message": "Something went wrong",
		})
		return
	}

	fmt.Println("refresh for user", uid)

	utils.ResponseWithMessage(ctx, http.StatusOK, gin.H{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	})
}
