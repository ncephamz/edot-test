package err

import "errors"

var (
	UnAuthorizedError = errors.New("Username or password not match")
	InvalidToken      = errors.New("invalid token")
)
