package domain

import (
	"time"
)

type UserRole string

const (
	Admin    UserRole = "Admin"
	Customer UserRole = "Customer"
)

type User struct {
	ID                 uint64
	Name               string
	Surname            string
	Email              string
	Password           string
	Address            string
	NotificationRadius uint64
	Role               UserRole
	CreatedAt          time.Time
	UpdatedAt          time.Time
}
