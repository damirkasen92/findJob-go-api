package model

import (
	"time"
)

type User struct {
	ID uint `gorm:"primaryKey"`

	Email    string `gorm:"unique"`
	Password string

	Role Role

	CreatedAt time.Time
	UpdatedAt time.Time
}
