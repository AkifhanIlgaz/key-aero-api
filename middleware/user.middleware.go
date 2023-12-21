package middleware

import (
	"github.com/AkifhanIlgaz/key-aero-api/services"
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
		// accessToken, err := utils.ParseAuthHeader(ctx)
		// if err != nil {
		// 	utils.ResponseWithMessage(ctx, http.StatusUnauthorized, gin.H{
		// 		"message": err.Error(),
		// 	})
		// 	return
		// }

		// claims, err := middleware.tokenService.ParseAccessToken(accessToken)
		// if err != nil {
		// 	utils.ResponseWithMessage(ctx, http.StatusUnauthorized, gin.H{
		// 		"message": err.Error(),
		// 	})
		// 	return
		// }

	}
}
