package port

import (
	"context"
	"github.com/bulutcan99/commerce_shipment/internal/core/domain"
)

type ITokenService interface {
	CreateToken(user *domain.User) (string, error)
	VerifyToken(token string) (*domain.TokenPayload, error)
}

type IAuthService interface {
	Login(ctx context.Context, email, password string) (string, error)
}
