package customErrors

import (
	"errors"
	"net/http"
)

var (
	ErrBadInputData = errors.New("bad input data")

	ErrUnauthorized  = errors.New("user is not authorized")
	ErrEmailNotValid = errors.New("this email is not valid")
	ErrEmailExist    = errors.New("this email already exists")
	ErrEmailNotExist = errors.New("this email doesn`t exists")
	ErrUserNotFound  = errors.New("this user is not found")
	ErrInvalidCode   = errors.New("invalid verification code")

	ErrShortPassword = errors.New("password should be longer than 6 characters")
	ErrLatinPassword = errors.New("password should contains Latin characters and numbers")

	ErrNoAccess = errors.New("user doesn't have access")

	ErrAlreadyAppended = errors.New("user has already been added")
)

var errorToCode = map[error]int{
	ErrBadInputData: http.StatusBadRequest,

	ErrUnauthorized:  http.StatusUnauthorized,
	ErrEmailNotValid: http.StatusBadRequest,
	ErrEmailExist:    http.StatusConflict,
	ErrEmailNotExist: http.StatusBadRequest,
	ErrUserNotFound:  http.StatusNotFound,
	ErrInvalidCode:   http.StatusBadRequest,

	ErrShortPassword: http.StatusBadRequest,
	ErrLatinPassword: http.StatusBadRequest,

	ErrAlreadyAppended: http.StatusConflict,

	ErrNoAccess: http.StatusForbidden,
}

func ConvertErrorToCode(err error) (code int) {
	code, isErrorExist := errorToCode[err]
	if !isErrorExist {
		code = http.StatusInternalServerError
	}
	return
}
