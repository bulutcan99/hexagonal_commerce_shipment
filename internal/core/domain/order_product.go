package domain

import (
	"time"
)

type OrderProduct struct {
	ID         uint64
	OrderID    uint64
	ProductID  uint64
	Quantitiy  int8
	TotalPrice float64
	Sale       int64
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Order      *Order
	Product    *Product
}
