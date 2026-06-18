package model

import (
	"time"

	"gorm.io/datatypes"
)

type Resume struct {
	ID uint

	Title  string
	About  string
	Skills datatypes.JSON `gorm:"type:jsonb"`
	UserID uint

	CreatedAt time.Time

	User User `gorm:"foreignKey:userID"`
}
