package paseto

import (
	"github.com/bulutcan99/commerce_shipment/internal/core/domain"
	"github.com/bulutcan99/commerce_shipment/internal/core/port"
	"time"

	"github.com/google/uuid"
	"github.com/o1egl/paseto"
	"golang.org/x/crypto/chacha20poly1305"
)

/**
 * PasetoToken implements port.TokenService interface
 * and provides an access to the paseto library
 */
type PasetoToken struct {
	paseto       *paseto.V2
	symmetricKey []byte
	duration     time.Duration
}

// New creates a new paseto instance
func New(config *config.Token) (port.ITokenService, error) {
	symmetricKey := config.SymmetricKey
	durationStr := config.Duration

	validSymmetricKey := len(symmetricKey) == chacha20poly1305.KeySize
	if !validSymmetricKey {
		return nil, domain.ErrInvalidTokenSymmetricKey
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

// CreateToken creates a new paseto token
func (pt *PasetoToken) CreateToken(user *domain.User) (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

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

// VerifyToken verifies the paseto token
func (pt *PasetoToken) VerifyToken(token string) (*domain.TokenPayload, error) {
	var payload domain.TokenPayload

	err := pt.paseto.Decrypt(token, pt.symmetricKey, &payload, nil)
	if err != nil {
		return nil, domain.ErrInvalidToken
	}

	isExpired := time.Now().After(payload.ExpiredAt)
	if isExpired {
		return nil, domain.ErrExpiredToken
	}

	return &payload, nil
}
