package domain

import (
	"errors"
)

// ErrorCode tipi, özel hata kodları için bir türdür.
type ErrorCode int

const (
	DataNotFound ErrorCode = iota + 1
	NoUpdatedData
	ConflictingData
	InsufficientStock
	InsufficientPayment
	ExpiredToken
	InvalidToken
	InvalidCredentials
	EmptyAuthorizationHeader
	InvalidAuthorizationHeader
	InvalidAuthorizationType
	Unauthorized
	Forbidden
	ErrInvalidEmail
	ErrInvalidRole
	ErrInvalidPassword
)

// Error struct, hata bilgilerini içerir.
type Error struct {
	Err     error
	Message string
	Code    ErrorCode
	Data    any
}

var (
	ErrDataNotFound               = errors.New("data not found")
	ErrNoUpdatedData              = errors.New("no data to update")
	ErrConflictingData            = errors.New("data conflicts with existing data in unique column")
	ErrInsufficientStock          = errors.New("product stock is not enough")
	ErrInsufficientPayment        = errors.New("total paid is less than total price")
	ErrExpiredToken               = errors.New("access token has expired")
	ErrInvalidToken               = errors.New("access token is invalid")
	ErrInvalidCredentials         = errors.New("invalid email or password")
	ErrEmptyAuthorizationHeader   = errors.New("authorization header is not provided")
	ErrInvalidAuthorizationHeader = errors.New("authorization header format is invalid")
	ErrInvalidAuthorizationType   = errors.New("authorization type is not supported")
	ErrUnauthorized               = errors.New("user is unauthorized to access the resource")
	ErrForbidden                  = errors.New("user is forbidden to access the resource")
)
