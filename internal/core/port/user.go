package port

import (
	"context"
)

type IUserRepository interface {
	CreateUser(ctx context.Context, user *_user.User) (*_user.User, error)
	GetUserByID(ctx context.Context, id uint64) (*_user.User, error)
	GetUserByEmail(ctx context.Context, email string) (*_user.User, error)
	ListUsers(ctx context.Context, skip, limit uint64) ([]_user.User, error)
	UpdateUser(ctx context.Context, user *_user.User) (*_user.User, error)
	DeleteUser(ctx context.Context, id uint64) error
}

type IUserService interface {
	Register(ctx context.Context, user *_user.User) (*_user.User, error)
	GetUserById(ctx context.Context, id uint64) (*_user.User, error)
	GetUserByEmail(ctx context.Context, email string) (*_user.User, error)
	ListUsers(ctx context.Context, skip, limit uint64) ([]_user.User, error)
	UpdateUser(ctx context.Context, user *_user.User) (*_user.User, error)
	DeleteUser(ctx context.Context, id uint64) error
}
