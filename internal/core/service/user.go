package service

import (
	"github.com/bulutcan99/commerce_shipment/internal/core/port"
)

type UserService struct {
	repo  port.IUserRepository
	cache port.ICacheRepository
}

func NewUserService(repo port.IUserRepository, cache port.ICacheRepository) *UserService {
	return &UserService{
		repo,
		cache,
	}
}
