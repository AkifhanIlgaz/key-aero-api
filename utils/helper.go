package utils

import (
	"fmt"

	"github.com/AkifhanIlgaz/key-aero-api/models"
	"github.com/gin-gonic/gin"
)

func GetUserFromContext(ctx *gin.Context) (*models.User, error) {
	val, exists := ctx.Get("user")

	if !exists {
		return nil, fmt.Errorf("user doesn't exist in gin context")
	}

	user, ok := val.(*models.User)
	if !ok {
		return nil, fmt.Errorf("invalid user type")
	}

	return user, nil
}

func ConvertToSliceString(anySlice []interface{}) []string {
	stringSlice := make([]string, len(anySlice))
	for i, v := range anySlice {
		stringSlice[i] = fmt.Sprint(v)
	}
	return stringSlice
}
