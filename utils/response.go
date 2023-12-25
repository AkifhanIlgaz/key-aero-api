package utils

import (
	"github.com/gin-gonic/gin"
)

var (
	successBase = gin.H{
		"status": "success",
	}
	failBase = gin.H{
		"status": "fail",
	}
)

func ResponseWithMessage(ctx *gin.Context, code int, data gin.H) {
	if code >= 200 && code < 300 {
		ctx.JSON(code, unzipData(successBase, data))
	} else {
		ctx.AbortWithStatusJSON(code, unzipData(failBase, data))
	}
}

func unzipData(base gin.H, data gin.H) gin.H {
	for key, value := range data {
		base[key] = value
	}
	return base
}
