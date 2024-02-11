package service

import (
	"context"
	"github.com/bulutcan99/commerce_shipment/internal/core/domain"
	"github.com/bulutcan99/commerce_shipment/internal/core/port"
	"github.com/bulutcan99/commerce_shipment/internal/core/util"
)

type UserService struct {
	userRepo port.IUserRepository
	cache    port.ICacheRepository
}

func NewUserService(userRepo port.IUserRepository, cache port.ICacheRepository) *UserService {
	return &UserService{
		userRepo,
		cache,
	}
}

func (us *UserService) ListAllUsers(ctx context.Context) ([]domain.User, *domain.Error) {
	var users []domain.User
	params := util.GenerateCacheKey("users", "all")
	cachedUsers, err := us.cache.Get(ctx, params)
	if err == nil {
		err := util.Deserialize(cachedUsers, &users)
		if err != nil {
			return nil, &domain.Error{
				Code:    domain.CacheSerialization,
				Message: "failed to deserialize users",
			}
		}
		return users, nil
	}

	users, errDomain := us.userRepo.GetAllUsers(ctx)
	if errDomain != nil {
		return nil, &domain.Error{
			Code:    domain.ErrSqlSelect,
			Message: "failed to get users from database",
		}
	}

	usersSerialized, err := util.Serialize(users)
	if err != nil {
		return nil, &domain.Error{
			Code:    domain.CacheSerialization,
			Message: "failed to serialize users",
		}
	}

	err = us.cache.Set(ctx, params, usersSerialized, 0)
	if err != nil {
		return nil, &domain.Error{
			Code:    domain.CacheSet,
			Message: "failed to set users in cache",
		}
	}

	return users, nil
}
