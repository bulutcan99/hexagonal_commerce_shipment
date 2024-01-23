package domain

import (
	"time"
)

type Order struct {
	ID           uint64
	UserID       uint64
	PaymentID    uint64
	CustomerName string
	TotalPrice   float64
	CreatedAt    time.Time
	UpdatedAt    time.Time
	User         *User
	Payment      *Payment
	Products     []OrderProduct
}
