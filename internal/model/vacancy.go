package model

import "time"

type Vacancy struct {
	ID          uint   `gorm:"primaryKey"`
	Title       string `gorm:"index"`
	Description string

	SalaryFrom int
	SalaryTo   int

	CreatedBy uint `gorm:"index"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
