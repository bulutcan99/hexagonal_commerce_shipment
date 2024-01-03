package port

import (
	"errors"
	"github.com/bulutcan99/shipment/internal/core/domain"
)

type ErrorFactory struct{}

func NewErrorFactory() *ErrorFactory {
	return &ErrorFactory{}
}

func (ef *ErrorFactory) NewError(code domain.ErrorCode, data interface{}) *domain.Error {
	var (
		err     error
		message string
	)

	switch code {
	case domain.DataNotFound:
		err = domain.ErrDataNotFound
		message = "Data not found"
	case domain.NoUpdatedData:
		err = domain.ErrNoUpdatedData
		message = "No data to update"
	case domain.ConflictingData:
		err = domain.ErrConflictingData
		message = "Data conflicts with existing data in unique column"
	case domain.InsufficientStock:
		err = domain.ErrInsufficientStock
		message = "Product stock is not enough"
	case domain.InsufficientPayment:
		err = domain.ErrInsufficientPayment
		message = "Total paid is less than total price"
	case domain.ExpiredToken:
		err = domain.ErrExpiredToken
		message = "Access token has expired"
	case domain.InvalidToken:
		err = domain.ErrInvalidToken
		message = "Access token is invalid"
	case domain.InvalidCredentials:
		err = domain.ErrInvalidCredentials
		message = "Invalid email or password"
	case domain.EmptyAuthorizationHeader:
		err = domain.ErrEmptyAuthorizationHeader
		message = "Authorization header is not provided"
	case domain.InvalidAuthorizationHeader:
		err = domain.ErrInvalidAuthorizationHeader
		message = "Authorization header format is invalid"
	case domain.InvalidAuthorizationType:
		err = domain.ErrInvalidAuthorizationType
		message = "Authorization type is not supported"
	case domain.Unauthorized:
		err = domain.ErrUnauthorized
		message = "User is unauthorized to access the resource"
	case domain.Forbidden:
		err = domain.ErrForbidden
		message = "User is forbidden to access the resource"
	default:
		err = errors.New("unknown error code")
		message = "Unknown error"
	}

	return &domain.Error{
		Err:     err,
		Message: message,
		Code:    int(code),
		Data:    data,
	}
}
