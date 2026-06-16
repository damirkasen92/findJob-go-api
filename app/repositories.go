package app

import (
	"github.com/damir/jobfinder/internal/repository"
	"gorm.io/gorm"
)

type Repositories struct {
	User        repository.UserRepository
	Vacancy     repository.VacancyRepository
	Resume      repository.ResumeRepository
	Application repository.ApplicationRepository
}

func NewRepositories(
	db *gorm.DB,
) *Repositories {
	return &Repositories{
		User:        repository.NewUserRepository(db),
		Vacancy:     repository.NewVacancyRepository(db),
		Resume:      repository.NewResumeRepository(db),
		Application: repository.NewApplicationRepository(db),
	}
}
