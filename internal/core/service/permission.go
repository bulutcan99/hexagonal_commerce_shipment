package service

import (
	"github.com/bulutcan99/commerce_shipment/internal/core/port"
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
