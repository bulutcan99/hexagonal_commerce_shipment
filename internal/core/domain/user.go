package domain

import (
	"time"
)

type User struct {
	ID                 uint64    `json:"id"`
	Name               string    `json:"name"`
	Surname            string    `json:"surname"`
	Email              string    `json:"email"`
	Password           string    `json:"password"`
	Address            string    `json:"address"`
	NotificationRadius uint64    `json:"notification_radius"`
	Role               string    `json:"role"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}
