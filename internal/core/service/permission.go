package service

import (
	"context"
	"github.com/bulutcan99/commerce_shipment/internal/core/domain"
	"github.com/bulutcan99/commerce_shipment/internal/core/port"
	"log/slog"
)

type PermissionService struct {
	repo  port.IPermissionRepository
	cache port.ICacheRepository
}

func NewPermissionService(repo port.IPermissionRepository, cache port.ICacheRepository) *PermissionService {
	return &PermissionService{
		repo,
		cache,
	}
}

func (p *PermissionService) AssignPermission(ctx context.Context, permission *domain.Permission, userId uint64) (*domain.Permission, *domain.Error) {
	insertPermission, errInsert := p.repo.Insert(ctx, permission, userId)
	if errInsert != nil {
		return nil, &domain.Error{
			Code:    errInsert.Code,
			Message: errInsert.Message,
			Data:    insertPermission,
		}
	}
	slog.Info("Permission Assigned to: %v", userId)
	return insertPermission, nil
}
