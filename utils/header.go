package utils

import (
	"strings"

	"github.com/AkifhanIlgaz/key-aero-api/errors"
	"github.com/gin-gonic/gin"
)

const (
	authHeaderKey = "Authorization"
	bearerScheme  = "Bearer"
)

// Parse authorization header and return the acces token
func ParseAuthHeader(ctx *gin.Context) (string, error) {
	authorizationHeader := strings.TrimSpace(ctx.Request.Header.Get(authHeaderKey))
	if authorizationHeader == "" {
		return "", errors.ErrAuthHeaderMissing
	}

	fields := strings.Fields(authorizationHeader)

	if len(fields) != 2 {
		return "", errors.ErrInvalidAuthScheme
	}

	if fields[0] != "Bearer" {
		return "", errors.ErrInvalidAuthScheme
	}

	return fields[1], nil
}
