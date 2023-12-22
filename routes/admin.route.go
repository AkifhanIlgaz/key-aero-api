package routes

import (
	"github.com/AkifhanIlgaz/key-aero-api/controllers"
	"github.com/AkifhanIlgaz/key-aero-api/middleware"
	"github.com/AkifhanIlgaz/key-aero-api/utils"
	"github.com/gin-gonic/gin"
)

type AdminRouteController struct {
	adminController *controllers.AdminController
	userMiddleware  *middleware.UserMiddleware
}

func NewAdminRouteController(adminController *controllers.AdminController, userMiddleware *middleware.UserMiddleware) *AdminRouteController {
	return &AdminRouteController{
		adminController: adminController,
		userMiddleware:  userMiddleware,
	}
}

func (routeController *AdminRouteController) AdminRoute(rg *gin.RouterGroup) {
	router := rg.Group("/admin/user", routeController.userMiddleware.ExtractUser(), routeController.userMiddleware.HasRole(utils.AdminRole))
	{
		router.GET("/all", routeController.adminController.GetAllUsers)
		router.POST("/add", routeController.adminController.AddUser)
		router.DELETE("/delete/:id", routeController.adminController.DeleteUser)
	}

}
