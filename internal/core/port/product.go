package port

import (
	"context"
	"github.com/bulutcan99/shipment/internal/core/domain"
)

type ProductRepository interface {
	CreateProduct(ctx context.Context, product *domain.Product) (*domain.Product, error)
	GetProductByID(ctx context.Context, id uint64) (*domain.Product, error)
	ListProducts(ctx context.Context, search string, categoryId, skip, limit uint64) ([]domain.Product, error)
	UpdateProduct(ctx context.Context, product *domain.Product) (*domain.Product, error)
	DeleteProduct(ctx context.Context, id uint64) error
}

type ProductService interface {
	CreateProduct(ctx context.Context, product *domain.Product) (*domain.Product, error)
	GetProduct(ctx context.Context, id uint64) (*domain.Product, error)
	ListProducts(ctx context.Context, search string, categoryId, skip, limit uint64) ([]domain.Product, error)
	UpdateProduct(ctx context.Context, product *domain.Product) (*domain.Product, error)
	DeleteProduct(ctx context.Context, id uint64) error
}
