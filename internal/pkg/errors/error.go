package errors

import (
	"github.com/go-kratos/kratos/v2/errors"
	"net/http"
)

var (
	ErrPasswordInvalid   = errors.BadRequest("Invalid password", "password is invalid")
	ErrUsernameInvalid   = errors.BadRequest("username is invalid", "username is invalid")
	ErrPhoneInvalid      = errors.BadRequest("phone is invalid", "phone is invalid")
	ErrEmailInvalid      = errors.BadRequest("email is invalid", "email is invalid")
	ErrUserNotFound      = errors.NotFound("user not found", "user not found")
	ErrUnauthenticated   = errors.Unauthorized(http.StatusText(http.StatusUnauthorized), http.StatusText(http.StatusUnauthorized))
	ErrInternalServer    = errors.InternalServer(http.StatusText(http.StatusInternalServerError), "Internal server error")
	ErrUserAlreadyExists = errors.Conflict("User already exists", "User already exists")
)
