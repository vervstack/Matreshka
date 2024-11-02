package user_errors

import (
	errors "github.com/Red-Sock/trace-errors"
	"google.golang.org/grpc/codes"
)

var (
	ErrValidation    = errors.NewUserError("Validation error", codes.InvalidArgument)
	ErrAlreadyExists = errors.NewUserError("Already exists", codes.AlreadyExists)
)
