package service

import (
	"context"

	"github.com/damir/jobfinder/internal/dto"
	"github.com/damir/jobfinder/internal/model"
	"github.com/damir/jobfinder/internal/repository"
)

type resumeService struct {
	repo repository.ResumeRepository
}

func NewResumeService(repo repository.ResumeRepository) ResumeService {
	return &resumeService{
		repo: repo,
	}
}

func (s *resumeService) Create(
	ctx context.Context,
	req dto.CreateResumeRequest,
	actor dto.Actor,
) error {
	resume := model.Resume{
		Title:  req.Title,
		About:  req.About,
		Skills: req.Skills,
		UserID: actor.UserID,
	}

	return s.repo.Create(
		ctx,
		&resume,
	)
}

func (s *resumeService) Delete(
	ctx context.Context,
	resumeID uint64,
	actor dto.Actor,
) error {
	resume, err := s.repo.GetByID(
		ctx,
		resumeID,
	)

	if err != nil {
		return err
	}

	if actor.Role == model.RoleAdmin || resume.UserID == actor.UserID {
		return s.repo.Delete(
			ctx,
			resumeID,
		)
	}

	return model.ErrForbidden
}

func (s *resumeService) GetByID(
	ctx context.Context,
	resumeID uint64,
) (*model.Resume, error) {
	resume, err := s.repo.GetByID(
		ctx,
		resumeID,
	)

	if err != nil {
		if err.Error() == "record not found" {
			return nil, model.ErrNotFound
		}

		return nil, err
	}

	return resume, nil
}

func (s *resumeService) GetList(
	ctx context.Context,
) ([]model.Resume, error) {
	resumes, err := s.repo.GetList(
		ctx,
	)

	if err != nil {
		return nil, err
	}

	return resumes, nil
}

func (s *resumeService) MyResumes(
	ctx context.Context,
	actor dto.Actor,
) ([]model.Resume, error) {
	resumes, err := s.repo.GetByUserID(
		ctx,
		uint64(actor.UserID),
	)

	if err != nil {
		return nil, err
	}

	return resumes, nil
}
