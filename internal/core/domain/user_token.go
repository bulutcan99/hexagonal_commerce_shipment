package domain

import (
	"github.com/oklog/ulid/v2"
	"time"
)

type TokenPayload struct {
	ID        ulid.ULID
	UserID    uint64
	Role      string
	IssuedAt  time.Time
	ExpiredAt time.Time
}
