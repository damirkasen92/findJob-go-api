package model

import "time"

type Application struct {
	ID uint `gorm:"primaryKey"`

	ResumeID  uint `gorm:"index:idx_resume_vacancy,unique"`
	VacancyID uint `gorm:"index:idx_resume_vacancy,unique"`

	UserID uint

	Status ApplicationStatus

	CreatedAt time.Time

	Resume Resume `gorm:"foreignKey:ResumeID"`
}
