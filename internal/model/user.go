package model

import "time"

type User struct {
	ID uint `gorm:"primaryKey"`

	Email    string `gorm:"unique"`
	Password string

	Role string

	CreatedAt time.Time
	UpdatedAt time.Time
}
