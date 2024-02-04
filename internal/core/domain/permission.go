package domain

import (
	"time"
)

// Permission Entry is a type for permission entry: 1-Customer, 2-Worker, 3-SuperAdmin
// AddFlag is a type for permission add flag
type Permission struct {
	ID         uint64
	Entry      int
	AddFlag    bool
	RemoveFlag bool
	AdminFlag  bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
