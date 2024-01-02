package shipment

import (
	"github/bulutcan99/shipment/internal/core/domain"
)

type ShipmentError struct {
	domain.Error
}

func NewShipmentError(err error, message string, code int, data interface{}) *ShipmentError {
	return &ShipmentError{
		Error: domain.Error{
			Err:     err,
			Message: message,
			Code:    code,
			Data:    data,
		},
	}
}
