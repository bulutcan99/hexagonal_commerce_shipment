package domain

import (
	"github.com/google/uuid"
	"time"
)

type TokenPayload struct {
	ID           uuid.UUID
	UserId       uint64
	PermissionId uint64
	IssuedAt     time.Time
	ExpiredAt    time.Time
}
