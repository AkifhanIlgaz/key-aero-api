package errors

import "errors"

var (
	ErrInvalidAuthScheme   error = errors.New("invalid auth scheme")          // 401
	ErrAuthHeaderMissing   error = errors.New("authorization header missing") // 401
	ErrSomethingWentWrong  error = errors.New("something went wrong")         // 500
	ErrUnexpectedMethod    error = errors.New("unexpected signing method")    // 400
	ErrInvalidToken        error = errors.New("invalid token")                // 401
	ErrRefreshTokenMissing error = errors.New("refresh token is missing")
)
