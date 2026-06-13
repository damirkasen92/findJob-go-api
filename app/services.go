package app

import (
	"github.com/damir/jobfinder/internal/auth"
	"github.com/damir/jobfinder/internal/service"
)

type Services struct {
	User    service.UserService
	Vacancy service.VacancyService
	Resume  service.ResumeService
}

func NewServices(
	repos *Repositories,
	jwt *auth.JWTManager,
) *Services {
	return &Services{
		User: service.NewUserService(
			repos.User,
			jwt,
		),
		Vacancy: service.NewVacancyService(
			repos.Vacancy,
		),
		Resume: service.NewResumeService(
			repos.Resume,
		),
	}
}
