package port

import (
	"context"
	"github/bulutcan99/shipment/internal/core/domain"
	"time"
)

type ITokenService interface {
	CreateToken(user *domain.User, duration time.Duration) (string, error)
	VerifyToken(token string) (*domain.TokenPayload, error)
}

type IAuthService interface {
	Login(ctx context.Context, email, password string) (string, error)
}
