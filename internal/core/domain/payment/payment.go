package payment

import (
	"time"
)

type PaymentType string

const (
	Cash PaymentType = "CASH"
	Card PaymentType = "CARD"
	Coin PaymentType = "COIN"
)

type Payment struct {
	ID        uint64      `json:"id"`
	Name      string      `json:"name"`
	Type      PaymentType `json:"type"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}
