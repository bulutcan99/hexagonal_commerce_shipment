package user

import (
	"time"
)

type UserRole string

const (
	Admin    UserRole = "admin"
	Customer UserRole = "customer"
)

type User struct {
	ID                 uint64            `json:"id"`
	Name               string            `json:"name"`
	Surname            string            `json:"surname"`
	Email              string            `json:"email"`
	Password           string            `json:"password"`
	Address            string            `json:"address"`
	NotificationRadius uint64            `json:"notification_radius"`
	Role               UserRole          `json:"role"`
	Friends            []UserInformation `json:"friends"`
	CreatedAt          time.Time         `json:"created_at"`
	UpdatedAt          time.Time         `json:"updated_at"`
}

type UserInformation struct {
	Name    string   `json:"name"`
	Surname string   `json:"surname"`
	Email   string   `json:"email"`
	Address string   `json:"address"`
	Role    UserRole `json:"role"`
}
