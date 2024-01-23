package port

import (
	"context"
	"github.com/bulutcan99/commerce_shipment/internal/core/domain"
)

type IUserRepository interface {
	Insert(ctx context.Context, user *domain.User) (*domain.User, *domain.Error)
	GetByID(ctx context.Context, id uint64) (*domain.User, *domain.Error)
	GetByEmail(ctx context.Context, email string) (*domain.User, *domain.Error)
	GetAll(ctx context.Context) ([]domain.User, *domain.Error)
	GetAllWithLimit(ctx context.Context, skip, limit uint64) ([]domain.User, *domain.Error)
	Update(ctx context.Context, user *domain.User) (*domain.User, *domain.Error)
	Delete(ctx context.Context, id uint64) *domain.Error
}

type IUserService interface {
	Register(ctx context.Context, user *domain.User) (*domain.User, error)
	GetByID(ctx context.Context, id uint64) (*domain.User, error)
	GetByEmail(ctx context.Context, email string) (*domain.User, error)
	ListUsers(ctx context.Context, skip, limit uint64) ([]domain.User, error)
	Update(ctx context.Context, user *domain.User) (*domain.User, error)
	Delete(ctx context.Context, id uint64) error
}
