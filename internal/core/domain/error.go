package domain

type ErrorCode int

const (
	DataNotFound ErrorCode = iota + 1
	NoUpdatedData
	ConflictingData
	InsufficientStock
	InsufficientPayment
	ExpiredToken
	PasswordHashing
	InvalidToken
	InvalidCredentials
	TokenCreation
	TokenVerification
	TokenExpired
	EmptyAuthorizationHeader
	InvalidAuthorizationHeader
	InvalidAuthorizationType
	Unauthorized
	Forbidden
	CacheSerialization
	CacheSet
	CacheDelete
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
