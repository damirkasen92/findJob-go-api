package service

import (
	"context"

	"github.com/damir/jobfinder/internal/dto"
	"github.com/damir/jobfinder/internal/model"
	"github.com/damir/jobfinder/internal/query"
	"github.com/damir/jobfinder/internal/repository"
)

type Actor struct {
	UserID uint
	Role   model.Role
}

type vacancyService struct {
	repo repository.VacancyRepository
}

func NewVacancyService(repo repository.VacancyRepository) VacancyService {
	return &vacancyService{
		repo: repo,
	}
}

func (s *vacancyService) Create(
	ctx context.Context,
	req dto.CreateVacancyRequest,
	userID uint,
) error {
	if req.SalaryTo < req.SalaryFrom {
		return model.ErrInvalidSalaryRange
	}

	vacancy := model.Vacancy{
		Title:       req.Title,
		Description: req.Description,
		SalaryFrom:  req.SalaryFrom,
		SalaryTo:    req.SalaryTo,
		CreatedBy:   userID,
	}

	return s.repo.Create(
		ctx,
		&vacancy,
	)
}

func (s *vacancyService) Delete(
	ctx context.Context,
	id uint,
	actor Actor,
) error {
	vacancy, err := s.repo.GetByID(ctx, id)

	if err != nil {
		return err
	}

	if actor.Role == model.RoleAdmin || actor.UserID == vacancy.CreatedBy {
		return s.repo.Delete(ctx, id)
	}

	return model.ErrForbidden
}

func (s *vacancyService) GetByID(
	ctx context.Context,
	id uint,
) (*model.Vacancy, error) {
	vacancy, err := s.repo.GetByID(ctx, id)

	if err != nil {
		if err.Error() == "record not found" {
			return nil, model.ErrNotFound
		}

		return nil, err
	}

	return vacancy, nil
}

func (s *vacancyService) List(
	ctx context.Context,
	filter query.VacancyFilter,
) ([]model.Vacancy, int64, error) {
	vacancies, total, err := s.repo.List(ctx, filter)

	if err != nil {
		return nil, 0, err
	}

	return vacancies, total, nil
}
