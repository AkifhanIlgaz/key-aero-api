package utils

import (
	"github.com/gin-gonic/gin"
)

const (
	statusSuccess = "success"
	statusFail    = "fail"
)

func ResponseWithMessage(ctx *gin.Context, code int, data gin.H) {
	ctx.JSON(code, gin.H{
		"status": responseMessage(code),
		"data":   data,
	})
}

func responseMessage(code int) string {
	switch {
	case code >= 200 && code < 300:
		return statusSuccess
	default:
		return statusFail
	}
}
