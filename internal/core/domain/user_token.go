package domain

import (
	"github.com/oklog/ulid/v2"
	"time"
)

type TokenPayload struct {
	ID        ulid.ULID `json:"id"`
	UserID    uint64    `json:"user_id"`
	Role      string    `json:"role"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}
