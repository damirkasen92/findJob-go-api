package service

import (
	"context"

	"github.com/damir/jobfinder/internal/dto"
	"github.com/damir/jobfinder/internal/model"
	"github.com/damir/jobfinder/internal/query"
)

type VacancyService interface {
	Create(
		ctx context.Context,
		req dto.CreateVacancyRequest,
		userID uint,
	) error

	Delete(
		ctx context.Context,
		id uint,
		actor Actor,
	) error

	GetByID(
		ctx context.Context,
		id uint,
	) (*model.Vacancy, error)

	List(
		ctx context.Context,
		filter query.VacancyFilter,
	) ([]model.Vacancy, int64, error)
}
