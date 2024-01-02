package controllers

import (
	"fmt"
	"net/http"

	"github.com/AkifhanIlgaz/key-aero-api/errors"
	"github.com/AkifhanIlgaz/key-aero-api/models"
	"github.com/AkifhanIlgaz/key-aero-api/services"
	"github.com/AkifhanIlgaz/key-aero-api/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slices"
)

type AdminController struct {
	userService  *services.UserService
	tokenService *services.TokenService
	emailService *services.EmailService
}

func NewAdminController(userService *services.UserService, tokenService *services.TokenService, emailService *services.EmailService) *AdminController {
	return &AdminController{
		userService:  userService,
		tokenService: tokenService,
		emailService: emailService,
	}
}

func (controller *AdminController) AddUser(ctx *gin.Context) {
	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		fmt.Println(err)
		utils.ResponseWithMessage(ctx, http.StatusBadRequest, gin.H{
			"message": "There are some missing required fields",
		})
		return
	}

	err := controller.userService.CreateUser(user)
	if err != nil {
		if errors.Is(err, errors.ErrUsernameTaken) {
			utils.ResponseWithMessage(ctx, http.StatusConflict, gin.H{
				"message": err.Error(),
			})
			return
		}
		utils.ResponseWithMessage(ctx, http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = controller.emailService.NewUser(user.Email, user.Username, user.Password)
	if err != nil {
		fmt.Println(err)
		utils.ResponseWithMessage(ctx, http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	utils.ResponseWithMessage(ctx, http.StatusCreated, gin.H{
		"message": "User successfully created!",
	})
}

func (controller *AdminController) GetAllUsers(ctx *gin.Context) {
	users, err := controller.userService.GetUsers()
	if err != nil {
		utils.ResponseWithMessage(ctx, http.StatusInternalServerError, nil)
		return
	}

	utils.ResponseWithMessage(ctx, http.StatusOK, gin.H{
		"users": users,
	})
}

func (controller *AdminController) UpdateUser(ctx *gin.Context) {
	var updates map[string]any

	if err := ctx.Bind(&updates); err != nil {
		utils.ResponseWithMessage(ctx, http.StatusBadRequest, gin.H{
			"message": "There are some missing required fields",
		})
		return
	}

	id := ctx.Params.ByName("id")
	if id == "" {
		utils.ResponseWithMessage(ctx, http.StatusBadRequest, gin.H{
			"message": "Query param id is missing",
		})
		return
	}

	if roles, ok := updates["roles"]; ok {
		if anySlice, ok := roles.([]any); ok {
			updates["roles"] = utils.GenerateRolesString(utils.ConvertToSliceString(anySlice))
		}
	}

	err := controller.userService.UpdateUser(id, updates)
	if err != nil {
		fmt.Println(err)
		utils.ResponseWithMessage(ctx, http.StatusInternalServerError, gin.H{
			"message": errors.ErrSomethingWentWrong.Error(),
		})
		return
	}

	utils.ResponseWithMessage(ctx, http.StatusOK, gin.H{
		"message": fmt.Sprintf("User %v is successfully updated", id),
	})
}

func (controller *AdminController) SearchUser(ctx *gin.Context) {
	var search models.SearchUserInput

	err := ctx.ShouldBindQuery(&search)
	if err != nil {
		utils.ResponseWithMessage(ctx, http.StatusBadRequest, gin.H{
			"message": "There are some missing required fields",
		})
		return
	}

	users, err := controller.userService.SearchUser(search)
	if err != nil {
		fmt.Println(err)
		return
	}

	utils.ResponseWithMessage(ctx, http.StatusOK, gin.H{
		"users": users,
	})
}

func (controller *AdminController) DeleteUser(ctx *gin.Context) {
	user, err := utils.GetUserFromContext(ctx)
	if err != nil {
		utils.ResponseWithMessage(ctx, http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
		return
	}

	toDeletedIds := ctx.QueryArray("id[]")
	if len(toDeletedIds) == 0 {
		utils.ResponseWithMessage(ctx, http.StatusBadRequest, gin.H{
			"message": "id param missing",
		})
		return
	}

	if slices.Contains(toDeletedIds, user.Id) {
		utils.ResponseWithMessage(ctx, http.StatusUnauthorized, gin.H{
			"message": "cannot delete yourself",
		})
		return
	}

	err = controller.userService.DeleteUser(toDeletedIds)
	if err != nil {
		utils.ResponseWithMessage(ctx, http.StatusInternalServerError, gin.H{
			"message": errors.ErrSomethingWentWrong.Error(),
		})
		return
	}

	utils.ResponseWithMessage(ctx, http.StatusOK, gin.H{
		"message": fmt.Sprintf("Users %v is successfully deleted", toDeletedIds),
	})
}
