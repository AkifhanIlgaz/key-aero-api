package controllers

import (
	"net/http"

	"github.com/AkifhanIlgaz/key-aero-api/errors"
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
	user, err := utils.GetUserFromContext(ctx)
	if err != nil {
		utils.ResponseWithMessage(ctx, http.StatusUnauthorized, gin.H{
			"message": errors.ErrNotLoggedIn.Error(),
		})
		return
	}

	utils.ResponseWithMessage(ctx, http.StatusOK, gin.H{
		"user": user,
	})
	return
}

func (controller *UserController) Update(ctx *gin.Context) {
	user, err := utils.GetUserFromContext(ctx)
	if err != nil {
		utils.ResponseWithMessage(ctx, http.StatusUnauthorized, gin.H{
			"message": errors.ErrNotLoggedIn.Error(),
		})
		return
	}

	var updates map[string]any

	if err := ctx.Bind(&updates); err != nil {
		utils.ResponseWithMessage(ctx, http.StatusBadRequest, gin.H{
			"message": "There are some missing required fields",
		})
		return
	}
	if roles, ok := updates["roles"]; ok {
		if anySlice, ok := roles.([]any); ok {
			updates["roles"] = utils.GenerateRolesString(utils.ConvertToSliceString(anySlice))
		}
	}

	err = controller.userService.UpdateUser(user.Id, updates)
	if err != nil {
		utils.ResponseWithMessage(ctx, http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	utils.ResponseWithMessage(ctx, http.StatusOK, gin.H{
		"message": "Successfully updated",
	})
}
