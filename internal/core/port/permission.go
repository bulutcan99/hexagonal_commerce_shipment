package port

import (
	"context"
	"github.com/bulutcan99/commerce_shipment/internal/core/domain"
)

type IPermissionRepository interface {
	Insert(ctx context.Context, permission *domain.Permission, userId uint64) (*domain.Permission, *domain.Error)
}

type IPermissionService interface {
	AssignPermission(ctx context.Context, permission *domain.Permission, userId uint64) (*domain.Permission, *domain.Error)
}
