package service

import (
	"context"
	"github.com/bulutcan99/commerce_shipment/internal/core/domain"
	"github.com/bulutcan99/commerce_shipment/internal/core/port"
	"github.com/bulutcan99/commerce_shipment/internal/core/util"
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

func (u *UserService) Register(ctx context.Context, user *domain.User) (*domain.User, *domain.Error) {
	hashedPassword, err := util.HashPassword(user.Password)
	if err != nil {
		return nil, &domain.Error{
			Code:    domain.PasswordHashing,
			Message: err.Error(),
			Data:    hashedPassword,
		}
	}
	user.Password = hashedPassword
	insertUser, errInsert := u.repo.Insert(ctx, user)
	if errInsert != nil {
		return nil, &domain.Error{
			Code:    errInsert.Code,
			Message: errInsert.Message,
			Data:    insertUser,
		}
	}

	cachingKey := util.GenerateCacheKey("user", user.ID)
	userSerialized, err := util.Serialize(user)
	if err != nil {
		return nil, &domain.Error{
			Code:    domain.CacheSerialization,
			Message: err.Error(),
			Data:    userSerialized,
		}
	}
	err = u.cache.Set(ctx, cachingKey, userSerialized, 0)
	if err != nil {
		return nil, &domain.Error{
			Code:    domain.CacheSet,
			Message: err.Error(),
			Data:    cachingKey,
		}
	}
	err = u.cache.DeleteByPrefix(ctx, "users:*") // delete all users cache because of new one
	if err != nil {
		return nil, &domain.Error{
			Code:    domain.CacheDelete,
			Message: err.Error(),
			Data:    cachingKey,
		}
	}
	return user, nil
}
