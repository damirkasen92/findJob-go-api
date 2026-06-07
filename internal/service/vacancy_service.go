package service

import (
	"context"

	"github.com/damir/jobfinder/internal/dto"
	"github.com/damir/jobfinder/internal/model"
	"github.com/damir/jobfinder/internal/repository"
)

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
