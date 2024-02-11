package paseto_token

import (
	"errors"
	"github.com/bulutcan99/commerce_shipment/internal/adapter/config"
	"github.com/bulutcan99/commerce_shipment/internal/core/domain"
	"github.com/bulutcan99/commerce_shipment/internal/core/port"
	"github.com/google/uuid"
	"time"

	"github.com/o1egl/paseto"
	"golang.org/x/crypto/chacha20poly1305"
)

type PasetoToken struct {
	paseto       *paseto.V2
	symmetricKey []byte
	TTL          time.Duration
}

func New(config *config.Token) (port.ITokenService, error) {
	symmetricKey := config.SymmetricKey
	durationStr := config.TTL

	validSymmetricKey := len(symmetricKey) == chacha20poly1305.KeySize
	if !validSymmetricKey {
		return nil, errors.New("invalid token symmetric key")
	}

	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		return nil, err
	}

	return &PasetoToken{
		paseto.NewV2(),
		[]byte(symmetricKey),
		duration,
	}, nil
}

func (pt *PasetoToken) CreateToken(user *domain.User) (string, *domain.Error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", &domain.Error{
			Code:    domain.TokenCreation,
			Message: "Token creation failed",
		}
	}

	payload := domain.TokenPayload{
		ID:           id,
		UserId:       user.ID,
		PermissionId: user.PermissionId,
		IssuedAt:     time.Now(),
		ExpiredAt:    time.Now().Add(pt.TTL),
	}

	token, err := pt.paseto.Encrypt(pt.symmetricKey, payload, nil)

	return token, nil

}

func (pt *PasetoToken) VerifyToken(token string) (*domain.TokenPayload, *domain.Error) {
	var payload domain.TokenPayload

	err := pt.paseto.Decrypt(token, pt.symmetricKey, &payload, nil)
	if err != nil {
		return nil, &domain.Error{
			Code:    domain.TokenVerification,
			Message: "Token verification failed",
		}
	}

	isExpired := time.Now().After(payload.ExpiredAt)
	if isExpired {
		return nil, &domain.Error{
			Code:    domain.TokenExpired,
			Message: "Token expired",
		}
	}

	return &payload, nil
}
