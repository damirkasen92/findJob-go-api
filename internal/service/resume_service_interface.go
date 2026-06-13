package service

import (
	"context"

	"github.com/damir/jobfinder/internal/dto"
	"github.com/damir/jobfinder/internal/model"
)

type ResumeService interface {
	Create(
		ctx context.Context,
		req dto.CreateResumeRequest,
		actor dto.Actor,
	) error

	Delete(
		ctx context.Context,
		resumeID uint64,
		actor dto.Actor,
	) error

	GetByID(
		ctx context.Context,
		resumeID uint64,
	) (*model.Resume, error)

	MyResumes(
		ctx context.Context,
		actor dto.Actor,
	) ([]model.Resume, error)

	GetList(
		ctx context.Context,
	) ([]model.Resume, error)
}
