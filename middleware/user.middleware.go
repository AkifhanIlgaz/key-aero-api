package middleware

import (
	"net/http"
	"strings"

	"github.com/AkifhanIlgaz/key-aero-api/errors"
	"github.com/AkifhanIlgaz/key-aero-api/models"
	"github.com/AkifhanIlgaz/key-aero-api/services"
	"github.com/AkifhanIlgaz/key-aero-api/utils"
	"github.com/gin-gonic/gin"
)

type UserMiddleware struct {
	userService  *services.UserService
	tokenService *services.TokenService
}

func NewUserMiddleware(userService *services.UserService, tokenService *services.TokenService) *UserMiddleware {
	return &UserMiddleware{
		userService:  userService,
		tokenService: tokenService,
	}
}

func (middleware *UserMiddleware) ExtractUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		accessToken, err := utils.ParseAuthHeader(ctx)
		if err != nil {
			utils.ResponseWithMessage(ctx, http.StatusUnauthorized, gin.H{
				"message": err.Error(),
			})
			return
		}

		claims, err := middleware.tokenService.ParseAccessToken(accessToken)
		if err != nil {
			utils.ResponseWithMessage(ctx, http.StatusUnauthorized, gin.H{
				"message": err.Error(),
			})
			return
		}

		user, err := middleware.userService.GetUserById(claims.Subject)
		if err != nil {
			utils.ResponseWithMessage(ctx, http.StatusUnauthorized, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.Set("user", user)
		ctx.Next()
	}
}

func (middleware *UserMiddleware) HasRole(role string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		val, exists := ctx.Get("user")
		if !exists {
			utils.ResponseWithMessage(ctx, http.StatusUnauthorized, gin.H{
				"message": errors.ErrNotLoggedIn.Error(),
			})
			return
		}

		user, ok := val.(*models.User)
		if !ok {
			utils.ResponseWithMessage(ctx, http.StatusUnauthorized, gin.H{
				"message": errors.ErrNotLoggedIn.Error(),
			})
			return
		}

		if !strings.Contains(user.Roles, role) {
			utils.ResponseWithMessage(ctx, http.StatusUnauthorized, gin.H{
				"message": "you do not have " + role + " role",
			})
			return
		}

		ctx.Next()
	}
}
