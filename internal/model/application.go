package model

import "time"

type Application struct {
	ID uint `gorm:"primaryKey"`

	ResumeID  uint
	VacancyID uint

	UserID uint

	CreatedAt time.Time
}
