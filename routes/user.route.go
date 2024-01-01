package routes

import (
	"github.com/AkifhanIlgaz/key-aero-api/controllers"
	"github.com/AkifhanIlgaz/key-aero-api/middleware"
	"github.com/gin-gonic/gin"
)

type UserRouteController struct {
	userController *controllers.UserController
	userMiddleware *middleware.UserMiddleware
}

func NewUserRouteController(userController *controllers.UserController, userMiddleware *middleware.UserMiddleware) *UserRouteController {
	return &UserRouteController{
		userController: userController,
		userMiddleware: userMiddleware,
	}
}

func (routeController *UserRouteController) UserRoute(rg *gin.RouterGroup) {
	router := rg.Group("/user", routeController.userMiddleware.ExtractUser())

	router.GET("/me", routeController.userController.Me)
	router.PUT("/me/update", routeController.userController.Update)
	router.PUT("/me/change-password", routeController.userController.ChangePassword)
	// TODO: Update password endpoint
}
