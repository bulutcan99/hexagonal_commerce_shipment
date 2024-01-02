package payment

import (
	"github/bulutcan99/shipment/internal/core/domain"
)

type PaymentError struct {
	domain.Error
}

func NewPaymentError(err error, message string, code int, data interface{}) *PaymentError {
	return &PaymentError{
		Error: domain.Error{
			Err:     err,
			Message: message,
			Code:    code,
			Data:    data,
		},
	}
}
