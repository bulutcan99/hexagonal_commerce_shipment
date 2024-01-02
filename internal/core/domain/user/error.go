package user

import (
	"github/bulutcan99/shipment/internal/core/domain"
)

type UserError struct {
	domain.Error
}

func NewUserError(err error, message string, code int, data interface{}) *UserError {
	return &UserError{
		Error: domain.Error{
			Err:     err,
			Message: message,
			Code:    code,
			Data:    data,
		},
	}
}
