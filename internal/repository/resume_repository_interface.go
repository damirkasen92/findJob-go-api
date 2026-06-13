package repository

import (
	"context"

	"github.com/damir/jobfinder/internal/model"
)

type ResumeRepository interface {
	Create(
		ctx context.Context,
		resume *model.Resume,
	) error

	Delete(
		ctx context.Context,
		resumeID uint64,
	) error

	GetByID(
		ctx context.Context,
		resumeID uint64,
	) (*model.Resume, error)

	GetByUserID(
		ctx context.Context,
		userID uint64,
	) ([]model.Resume, error)

	GetList(
		ctx context.Context,
	) ([]model.Resume, error)
}
