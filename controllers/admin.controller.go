package controllers

import (
	"fmt"
	"net/http"

	"github.com/AkifhanIlgaz/key-aero-api/errors"
	"github.com/AkifhanIlgaz/key-aero-api/models"
	"github.com/AkifhanIlgaz/key-aero-api/services"
	"github.com/AkifhanIlgaz/key-aero-api/utils"
	"github.com/gin-gonic/gin"
)

type AdminController struct {
	userService  *services.UserService
	tokenService *services.TokenService
}

func NewAdminController(userService *services.UserService, tokenService *services.TokenService) *AdminController {
	return &AdminController{
		userService:  userService,
		tokenService: tokenService,
	}
}

func (controller *AdminController) AddUser(ctx *gin.Context) {
	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
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
	var user models.UpdateInput

	if err := ctx.Bind(&user); err != nil {
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
	user.Id = id

	err := controller.userService.UpdateUser(&user)
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

// TODO: Implement this function
func (controller *AdminController) SearchUser(ctx *gin.Context) {
	var search models.SearchInput

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
	// username ?
}

func (controller *AdminController) DeleteUser(ctx *gin.Context) {
	user, err := utils.GetUserFromContext(ctx)
	if err != nil {
		utils.ResponseWithMessage(ctx, http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
		return
	}

	toDeletedId := ctx.Params.ByName("id")
	if toDeletedId == "" {
		utils.ResponseWithMessage(ctx, http.StatusBadRequest, gin.H{
			"message": "id param missing",
		})
		return
	}

	if user.Id == toDeletedId {
		utils.ResponseWithMessage(ctx, http.StatusUnauthorized, gin.H{
			"message": "cannot delete yourself",
		})
		return
	}

	err = controller.userService.DeleteUser(toDeletedId)
	if err != nil {
		utils.ResponseWithMessage(ctx, http.StatusInternalServerError, gin.H{
			"message": errors.ErrSomethingWentWrong.Error(),
		})
		return
	}

	utils.ResponseWithMessage(ctx, http.StatusOK, gin.H{
		"message": fmt.Sprintf("User %v is successfully deleted", toDeletedId),
	})
}
