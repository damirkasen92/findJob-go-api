package service

import (
	"context"
	"fmt"
	"strings"

	"github.com/damir/jobfinder/internal/dto"
	"github.com/damir/jobfinder/internal/model"
	"github.com/damir/jobfinder/internal/repository"
	"gorm.io/datatypes"
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
	skills := strings.Split(req.Skills, ",")

	for i := range skills {
		skills[i] = strings.TrimSpace(skills[i])
	}

	resume := model.Resume{
		Title:  req.Title,
		About:  req.About,
		Skills: datatypes.JSON([]byte(fmt.Sprintf(`["%s"]`, strings.Join(skills, `","`)))),
		UserID: actor.UserID,
	}

	return s.repo.Create(
		ctx,
		&resume,
	)
}

func (s *resumeService) Update(
	ctx context.Context,
	dto dto.UpdateResumeRequest,
	actor dto.Actor,
) error {
	resume, err := s.repo.GetByID(ctx, dto.Id)

	if err != nil {
		return err
	}

	if actor.Role == model.RoleAdmin || resume.UserID == actor.UserID {
		return s.repo.Update(
			ctx,
			dto,
		)
	}

	return model.ErrForbidden
}

func (s *resumeService) Delete(
	ctx context.Context,
	resumeID uint,
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
	resumeID uint,
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
		actor.UserID,
	)

	if err != nil {
		return nil, err
	}

	return resumes, nil
}
