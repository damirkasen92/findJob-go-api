package service

import (
	"context"

	"github.com/damir/jobfinder/internal/dto"
	"github.com/damir/jobfinder/internal/model"
	"github.com/damir/jobfinder/internal/repository"
)

type applicationService struct {
	repo        repository.ApplicationRepository
	vacancyRepo repository.VacancyRepository
	resumeRepo  repository.ResumeRepository
}

func NewApplicationService(
	repo repository.ApplicationRepository,
	vacancyRepo repository.VacancyRepository,
	resumeRepo repository.ResumeRepository,
) ApplicationService {
	return &applicationService{
		repo:        repo,
		vacancyRepo: vacancyRepo,
		resumeRepo:  resumeRepo,
	}
}

func (s *applicationService) Create(
	ctx context.Context,
	req dto.CreateApplicationRequest,
	actor dto.Actor,
) error {
	isExist, err := s.Exists(ctx, req.ResumeID, req.VacancyID)

	if err != nil {
		return err
	}

	if isExist {
		return model.ErrAlreadyApplied
	}

	vacancy, err := s.vacancyRepo.GetByID(ctx, req.VacancyID)

	if err != nil {
		return err
	}

	if vacancy.CreatedBy == actor.UserID {
		return model.ErrOwnVacancy
	}

	resume, err := s.resumeRepo.GetByID(ctx, req.ResumeID)

	if err != nil {
		return err
	}

	if resume.UserID != actor.UserID {
		return model.ErrForeignResume
	}

	application :=
		model.Application{
			ResumeID:  req.ResumeID,
			VacancyID: req.VacancyID,
			UserID:    actor.UserID,
		}

	err = s.repo.Create(ctx, &application)

	if err != nil {
		return err
	}

	return nil
}

func (s *applicationService) Exists(
	ctx context.Context,
	resumeID uint,
	vacancyID uint,
) (bool, error) {
	isExist, err := s.repo.Exists(ctx, resumeID, vacancyID)

	if err != nil {
		return false, err
	}

	return isExist, nil
}

func (s *applicationService) GetByID(
	ctx context.Context,
	applicationID uint,
	actor dto.Actor,
) (*model.Application, error) {
	application, err := s.repo.GetByID(
		ctx,
		applicationID,
	)

	if err != nil {
		return nil, err
	}

	if application.UserID != actor.UserID {
		return nil, model.ErrForeignResume
	}

	return application, nil
}

func (s *applicationService) ListByVacancy(
	ctx context.Context,
	vacancyID uint,
) ([]model.Application, error) {
	applications, err := s.repo.ListByVacancy(
		ctx,
		vacancyID,
	)

	if err != nil {
		return nil, err
	}

	return applications, nil
}

func (s *applicationService) ListByUser(
	ctx context.Context,
	userID uint,
) ([]model.Application, error) {
	applications, err := s.repo.ListByUser(
		ctx,
		userID,
	)

	if err != nil {
		return nil, err
	}

	return applications, nil
}
