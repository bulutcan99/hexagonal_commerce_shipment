package port

import (
	"context"
	"github.com/bulutcan99/commerce_shipment/internal/core/domain"
)

type IUserRepository interface {
	AddUser(ctx context.Context, user *domain.User, permission *domain.Permission) (*domain.User, *domain.Error)
	GetUserByID(ctx context.Context, id uint64) (*domain.User, *domain.Error)
	GetUserByEmail(ctx context.Context, email string) (*domain.User, *domain.Error)
	GetAllUsers(ctx context.Context) ([]domain.User, *domain.Error)
	GetAllUsersWithLimit(ctx context.Context, skip, limit uint64) ([]domain.User, *domain.Error)
	UpdateUser(ctx context.Context, user *domain.User) (*domain.User, *domain.Error)
	DeleteUser(ctx context.Context, id uint64) *domain.Error
}

type IUserService interface {
	ListAllUsers(ctx context.Context) ([]domain.User, *domain.Error)
	ListUsersWithLimit(ctx context.Context, skip, limit uint64) ([]domain.User, *domain.Error)
	// GetByID(ctx context.Context, id uint64) (*domain.User, error)
	// Update(ctx context.Context, user *domain.User) (*domain.User, error)
	// GetByEmail(ctx context.Context, email string) (*domain.User, error)
	// Delete(ctx context.Context, id uint64) error
}
