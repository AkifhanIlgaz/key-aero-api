package routes

import (
	"github.com/AkifhanIlgaz/key-aero-api/controllers"
	"github.com/gin-gonic/gin"
)

type AuthRouteController struct {
	authController *controllers.AuthController
}

func NewAuthRouteController(authController *controllers.AuthController) *AuthRouteController {
	return &AuthRouteController{
		authController: authController,
	}
}

func (routeController *AuthRouteController) AuthRoute(rg *gin.RouterGroup) {
	router := rg.Group("/auth")
	// TODO: Middlewares ?

	router.POST("/signin", routeController.authController.SignIn)
	// TODO: Implement
	// ? Extract user middleware
	router.POST("/signout", routeController.authController.SignOut)

	router.POST("/refresh", routeController.authController.Refresh)
}
