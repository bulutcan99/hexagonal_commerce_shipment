package port

import (
	"context"
	"github.com/bulutcan99/commerce_shipment/internal/core/domain"
)

type IPermissionRepository interface {
	AddPermission(ctx context.Context, permission *domain.Permission) (*domain.Permission, *domain.Error)
	// Update(ctx context.Context, permission *domain.Permission) (*domain.Permission, *domain.Error)
}

type IPermissionService interface {
	// UpdatePermission(ctx context.Context, permission *domain.Permission, userId uint64) (*domain.Permission, *domain.Error)
}
