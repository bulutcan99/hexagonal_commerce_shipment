package port

import (
	"context"
	"github.com/bulutcan99/commerce_shipment/internal/core/domain"
)

type ITokenService interface {
	CreateToken(user *domain.User) (string, *domain.Error)
	VerifyToken(token string) (*domain.TokenPayload, *domain.Error)
}

type IAuthService interface {
	Register(ctx context.Context, user *domain.User) (*domain.User, *domain.Error)
	Login(ctx context.Context, email, password string) (string, *domain.Error)
}
