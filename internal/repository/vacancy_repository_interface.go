package repository

import (
	"context"

	"github.com/damir/jobfinder/internal/model"
	"github.com/damir/jobfinder/internal/query"
)

type VacancyRepository interface {
	Create(
		ctx context.Context,
		vacancy *model.Vacancy,
	) error

	GetByID(
		ctx context.Context,
		id uint,
	) (*model.Vacancy, error)

	Update(
		ctx context.Context,
		vacancy *model.Vacancy,
	) error

	Delete(
		ctx context.Context,
		id uint,
	) error

	List(
		ctx context.Context,
		filter query.VacancyFilter,
	) ([]model.Vacancy, error)
}
