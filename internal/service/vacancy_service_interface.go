package service

import (
	"context"

	"github.com/damir/jobfinder/internal/dto"
)

type VacancyService interface {
	Create(
		ctx context.Context,
		req dto.CreateVacancyRequest,
		userID uint,
	) error
}
