package domain

import (
	"time"
)

type Shipment struct {
	ID              uint64
	DriverName      string
	ProductId       uint64
	CustomerName    string
	CustomerAddress string
	CustomerRadius  uint64
	Lat             float64
	Lng             float64
	Address         string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
