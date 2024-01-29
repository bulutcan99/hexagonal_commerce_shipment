package domain

import (
	"time"
)

// Entry is a type for permission entry

type Permission struct {
	ID        uint64
	Entry     int
	AddFlag   bool
	AdminFlag bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
