package domain

import (
	"github.com/oklog/ulid/v2"
	"time"
)

type Product struct {
	ID         uint64
	CategoryID uint64
	SKU        ulid.ULID
	Name       string
	Stock      int64
	Price      float64
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Category   *Category
}
