package product

import (
	"github.com/oklog/ulid/v2"
	"time"
)

type Product struct {
	ID         uint64    `json:"id"`
	CategoryID uint64    `json:"category_id"`
	SKU        ulid.ULID `json:"sku"` // Sku means Stock Keeping Unit
	Name       string    `json:"name"`
	Stock      int64     `json:"stock"`
	Price      float64   `json:"price"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Category   *Category `json:"category"`
}
