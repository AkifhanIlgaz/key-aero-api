package utils

import (
	"github.com/gin-gonic/gin"
)

const (
	statusSuccess = "success"
	statusFail    = "fail"
)

func ResponseWithMessage(ctx *gin.Context, code int, data gin.H) {
	if code >= 200 && code < 300 {
		ctx.JSON(code, gin.H{
			"status": statusSuccess,
			"data":   data,
		})
	} else {
		ctx.AbortWithStatusJSON(code, gin.H{
			"status": statusFail,
			"data":   data,
		})
	}

}
