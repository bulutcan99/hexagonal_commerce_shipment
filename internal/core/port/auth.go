package port

import (
	"context"
	"github/bulutcan99/shipment/internal/core/domain/session"
	"time"
)

type ITokenService interface {
	CreateToken(user *_user.User, duration time.Duration) (string, error)
	VerifyToken(token string) (*session.TokenPayload, error)
}

type IAuthService interface {
	Login(ctx context.Context, email, password string) (string, error)
}
