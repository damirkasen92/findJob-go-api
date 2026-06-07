package model

import "time"

type Vacancy struct {
	ID          uint `gorm:"primaryKey"`
	Title       string
	Description string

	SalaryFrom int
	SalaryTo   int

	CreatedBy uint

	CreatedAt time.Time
	UpdatedAt time.Time
}
