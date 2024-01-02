package product

import (
	"github/bulutcan99/shipment/internal/core/domain"
)

type ProductError struct {
	domain.Error
}

func NewProductError(err error, message string, code int, data interface{}) *ProductError {
	return &ProductError{
		Error: domain.Error{
			Err:     err,
			Message: message,
			Code:    code,
			Data:    data,
		},
	}
}
