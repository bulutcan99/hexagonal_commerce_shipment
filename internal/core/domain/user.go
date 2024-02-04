package domain

import (
	"time"
)

type User struct {
	ID                 uint64
	Name               string
	Surname            string
	Email              string
	Password           string
	Address            string
	NotificationRadius uint64
	Permission         Permission
	CreatedAt          time.Time
	UpdatedAt          time.Time
}
