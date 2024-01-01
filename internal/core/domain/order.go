package domain

import (
	"time"
)

type Order struct {
	ID           uint64         `json:"id"`
	UserID       uint64         `json:"user_id"`
	PaymentID    uint64         `json:"payment_id"`
	CustomerName string         `json:"customer_name"`
	TotalPrice   float64        `json:"total_price"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	User         *User          `json:"user"`
	Payment      *Payment       `json:"payment"`
	Products     []OrderProduct `json:"products"`
}
