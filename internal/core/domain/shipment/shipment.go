package shipment

import (
	"time"
)

type Shipment struct {
	ID              uint64    `json:"id"`
	DriverName      string    `json:"driver_name"`
	ProductId       uint64    `json:"product"`
	CustomerName    string    `json:"customer_name"`
	CustomerAddress string    `json:"customer_address"`
	CustomerRadius  uint64    `json:"notification_radius"`
	Lat             float64   `json:"lat"`
	Lng             float64   `json:"lng"`
	Address         string    `json:"address"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
