package repository

import (
	"context"

	"github.com/damir/jobfinder/internal/model"
)

type ApplicationRepository interface {
	Create(
		ctx context.Context,
		application *model.Application,
	) error

	GetByID(
		ctx context.Context,
		applicationID uint,
	) (*model.Application, error)

	Exists(
		ctx context.Context,
		resumeID uint,
		vacancyID uint,
	) (bool, error)

	ListByVacancy(
		ctx context.Context,
		vacancyID uint,
	) ([]model.Application, error)

	ListByUser(
		ctx context.Context,
		userID uint,
	) ([]model.Application, error)
}
