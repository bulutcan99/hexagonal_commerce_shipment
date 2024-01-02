package order

import (
	"github/bulutcan99/shipment/internal/core/domain"
)

type OrderError struct {
	domain.Error
}

func NewOrderError(err error, message string, code int, data interface{}) *OrderError {
	return &OrderError{
		Error: domain.Error{
			Err:     err,
			Message: message,
			Code:    code,
			Data:    data,
		},
	}
}
