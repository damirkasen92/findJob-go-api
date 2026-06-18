package service

import (
	"context"

	"github.com/damir/jobfinder/internal/dto"
	"github.com/damir/jobfinder/internal/model"
)

type ApplicationService interface {
	Create(
		ctx context.Context,
		req dto.CreateApplicationRequest,
		actor dto.Actor,
	) error

	UpdateStatus(
		ctx context.Context,
		appId uint,
		dto dto.UpdateApplicationStatusRequest,
	) error

	Exists(
		ctx context.Context,
		resumeID uint,
		vacancyID uint,
	) (bool, error)

	GetByID(
		ctx context.Context,
		applicationID uint,
		actor dto.Actor,
	) (*model.Application, error)

	ListByVacancy(
		ctx context.Context,
		vacancyID uint,
		actor dto.Actor,
	) ([]model.Application, error)

	ListByUser(
		ctx context.Context,
		actor dto.Actor,
	) ([]model.Application, error)
}
