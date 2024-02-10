package service

import (
	"github.com/bulutcan99/commerce_shipment/internal/core/port"
)

type UserService struct {
	userRepo port.IUserRepository
	permRepo port.IPermissionRepository
	cache    port.ICacheRepository
}

func NewUserService(userRepo port.IUserRepository, permRepo port.IPermissionRepository, cache port.ICacheRepository) *UserService {
	return &UserService{
		userRepo,
		permRepo,
		cache,
	}
}
