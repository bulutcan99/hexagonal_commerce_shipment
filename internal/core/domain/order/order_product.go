package order

import (
	"github/bulutcan99/shipment/internal/core/domain/product"
	"time"
)

type OrderProduct struct {
	ID         uint64           `json:"id"`
	OrderID    uint64           `json:"order_id"`
	ProductID  uint64           `json:"product_id"`
	Quantitiy  int8             `json:"quantitiy"`
	TotalPrice float64          `json:"total_price"`
	Sale       int64            `json:"sale"`
	CreatedAt  time.Time        `json:"created_at"`
	UpdatedAt  time.Time        `json:"updated_at"`
	Order      *Order           `json:"order"`
	Product    *product.Product `json:"product"`
}
