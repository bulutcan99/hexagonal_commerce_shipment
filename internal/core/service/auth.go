package service

import (
	"context"
	"github.com/bulutcan99/commerce_shipment/internal/core/domain"
	"github.com/bulutcan99/commerce_shipment/internal/core/port"
	"github.com/bulutcan99/commerce_shipment/internal/core/util"
)

type AuthService struct {
	userRepo port.IUserRepository
	permRepo port.IPermissionRepository
	cache    port.ICacheRepository
	token    port.ITokenService
}

func NewAuthService(userRepo port.IUserRepository, permRepo port.IPermissionRepository, cache port.ICacheRepository, token port.ITokenService) *AuthService {
	return &AuthService{
		userRepo,
		permRepo,
		cache,
		token,
	}
}

func (as *AuthService) Register(ctx context.Context, user *domain.User) (*domain.User, *domain.Error) {
	defaultPermission := domain.Permission{
		Entry:      1,
		AddFlag:    true,
		RemoveFlag: false,
		AdminFlag:  false,
	}

	insertPermission, errInsert := as.permRepo.Insert(ctx, &defaultPermission)

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
	insertUser, errInsert := as.userRepo.Insert(ctx, user, insertPermission)
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
	err = as.cache.Set(ctx, cachingKey, userSerialized, 0)
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

	err = as.cache.Set(ctx, permissionKey, permissionSerialized, 0)
	err = as.cache.DeleteByPrefix(ctx, "users:*") // delete all users cache because of new one
	if err != nil {
		return nil, &domain.Error{
			Code:    domain.CacheDelete,
			Message: err.Error(),
			Data:    cachingKey,
		}
	}
	return user, nil
}

func (as *AuthService) Login(ctx context.Context, email, password string) (string, *domain.Error) {
	user, err := as.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return "", &domain.Error{
			Code:    domain.InvalidCredentials,
			Message: "Invalid credentials",
			Data:    email,
		}
	}

	passErr := util.ComparePassword(password, user.Password)
	if passErr != nil {
		return "", &domain.Error{
			Code:    domain.InvalidCredentials,
			Message: "Invalid credentials",
			Data:    user.Password,
		}
	}

	token, err := as.token.CreateToken(user)
	if err != nil {
		return "", &domain.Error{
			Code:    domain.TokenCreation,
			Message: "Token creation failed",
		}
	}

	return token, nil
}
