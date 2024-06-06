package errs

import (
	"net/http"
)

type (
	// ValidationError описывает ошибки валидации
	ValidationError struct {
		BaseError
	}
	// DomainError описывает ошибки бизнес-логики
	DomainError struct {
		BaseError
	}
	// SystemError описывает системные ошибки
	SystemError struct {
		BaseError
	}
)

func NewValidationError(msg string, err error) error {
	return &ValidationError{
		BaseError{
			msg:  msg,
			err:  err,
			code: http.StatusBadRequest,
		},
	}
}

func NewDomainError(msg string, err error) error {
	return &DomainError{
		BaseError{
			msg:  msg,
			err:  err,
			code: http.StatusConflict,
		},
	}
}

func NewSystemError(msg string, err error) error {
	return &SystemError{
		BaseError{
			msg:  msg,
			err:  err,
			code: http.StatusInternalServerError,
		},
	}
}
