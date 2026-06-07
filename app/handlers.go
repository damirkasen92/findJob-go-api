package app

import "github.com/damir/jobfinder/internal/handler"

type Handlers struct {
	Auth    *handler.AuthHandler
	Vacancy *handler.VacancyHandler
}

func NewHandlers(
	services *Services,
) *Handlers {
	return &Handlers{
		Auth:    handler.NewAuthHandler(services.User),
		Vacancy: handler.NewVacancyHandler(services.Vacancy),
	}
}
