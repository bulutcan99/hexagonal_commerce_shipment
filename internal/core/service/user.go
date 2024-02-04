package service

import (
	"context"
	"github.com/bulutcan99/commerce_shipment/internal/core/domain"
	"github.com/bulutcan99/commerce_shipment/internal/core/port"
	"github.com/bulutcan99/commerce_shipment/internal/core/util"
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

func (u *UserService) Register(ctx context.Context, user *domain.User) (*domain.User, *domain.Error) {
	defaultPermission := domain.Permission{
		Entry:      1,
		AddFlag:    true,
		RemoveFlag: false,
		AdminFlag:  false,
	}

	insertPermission, errInsert := u.permRepo.Insert(ctx, &defaultPermission)

	if errInsert != nil {
		return nil, &domain.Error{
			Code:    errInsert.Code,
			Message: errInsert.Message,
			Data:    insertPermission.ID,
		}
	}

	hashedPassword, err := util.HashPassword(user.Password)
	if err != nil {
		return nil, &domain.Error{
			Code:    domain.PasswordHashing,
			Message: err.Error(),
			Data:    hashedPassword,
		}
	}
	user.Password = hashedPassword
	insertUser, errInsert := u.userRepo.Insert(ctx, user, insertPermission)
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

	permissionKey := util.GenerateCacheKeyParams("user", "permission", insertPermission.ID)
	permissionSerialized, err := util.Serialize(insertPermission)
	if err != nil {
		return nil, &domain.Error{
			Code:    domain.CacheSerialization,
			Message: err.Error(),
			Data:    permissionSerialized,
		}
	}

	err = u.cache.Set(ctx, permissionKey, permissionSerialized, 0)
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
