package paseto

import (
	"errors"
	"github.com/bulutcan99/commerce_shipment/internal/adapter/config"
	"github.com/bulutcan99/commerce_shipment/internal/core/domain"
	"github.com/bulutcan99/commerce_shipment/internal/core/port"
	"github.com/oklog/ulid/v2"
	"time"

	"github.com/o1egl/paseto"
	"golang.org/x/crypto/chacha20poly1305"
)

type PasetoToken struct {
	paseto       *paseto.V2
	symmetricKey []byte
	duration     time.Duration
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

func (pt *PasetoToken) CreateToken(user *domain.User) (string, error) {
	id := ulid.Make()

	payload := domain.TokenPayload{
		ID:        id,
		UserID:    user.ID,
		Role:      user.Role,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(pt.duration),
	}

	token, err := pt.paseto.Encrypt(pt.symmetricKey, payload, nil)

	return token, err

}

func (pt *PasetoToken) VerifyToken(token string) (*domain.TokenPayload, error) {
	var payload domain.TokenPayload

	err := pt.paseto.Decrypt(token, pt.symmetricKey, &payload, nil)
	if err != nil {
		return nil, errors.New("invalid token")
	}

	isExpired := time.Now().After(payload.ExpiredAt)
	if isExpired {
		return nil, errors.New("token is expired")
	}

	return &payload, nil
}
