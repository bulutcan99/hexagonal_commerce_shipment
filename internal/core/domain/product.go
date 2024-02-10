package domain

import (
	"github.com/google/uuid"
	"time"
)

type Product struct {
	ID         uint64
	CategoryID uint64
	SKU        uuid.UUID
	Name       string
	Stock      int64
	Price      float64
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Category   *Category
}
