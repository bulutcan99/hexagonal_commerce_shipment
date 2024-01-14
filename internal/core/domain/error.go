package domain

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
	ErrSqlInsert
	ErrSqlSelect
	ErrSqlDelete
	ErrSqlUpdate
)

type Error struct {
	Err     error
	Message string
	Code    ErrorCode
	Data    any
}
