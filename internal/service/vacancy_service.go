package service

import (
	"context"
	"errors"

	"github.com/damir/jobfinder/internal/dto"
	"github.com/damir/jobfinder/internal/model"
	"github.com/damir/jobfinder/internal/repository"
)

var ErrInvalidSalaryRange = errors.New("invalid salary range")

type Actor struct {
	UserID uint
	Role   model.Role
}

var ErrForbidden = errors.New("forbidden")

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
		return ErrInvalidSalaryRange
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

	return ErrForbidden
}

func (s *vacancyService) GetByID(
	ctx context.Context,
	id uint,
) (*model.Vacancy, error) {
	vacancy, err := s.repo.GetByID(ctx, id)

	if err != nil {
		return nil, err
	}

	return vacancy, nil
}

func (s *vacancyService) List(
	ctx context.Context,
) ([]model.Vacancy, error) {
	vacancies, err := s.repo.List(ctx)

	if err != nil {
		return nil, err
	}

	return vacancies, nil
}
