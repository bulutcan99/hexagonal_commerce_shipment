package domain

import (
	"errors"
)

// ErrorCode tipi, özel hata kodları için bir türdür.
type ErrorCode int

const (
	DataNotFound               ErrorCode = iota + 1
	NoUpdatedData              ErrorCode = iota + 2
	ConflictingData            ErrorCode = iota + 3
	InsufficientStock          ErrorCode = iota + 4
	InsufficientPayment        ErrorCode = iota + 5
	ExpiredToken               ErrorCode = iota + 6
	InvalidToken               ErrorCode = iota + 7
	InvalidCredentials         ErrorCode = iota + 8
	EmptyAuthorizationHeader   ErrorCode = iota + 9
	InvalidAuthorizationHeader ErrorCode = iota + 10
	InvalidAuthorizationType   ErrorCode = iota + 11
	Unauthorized               ErrorCode = iota + 12
	Forbidden                  ErrorCode = iota + 13
)

// Error struct, hata bilgilerini içerir.
type Error struct {
	Err     error
	Message string
	Code    int
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
