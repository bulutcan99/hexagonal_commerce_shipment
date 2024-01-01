package domain

import (
	"time"
)

type UserRole string

const (
	Admin    UserRole = "admin"
	Customer UserRole = "customer"
)

type User struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Address   string    `json:"address"`
	Role      UserRole  `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
