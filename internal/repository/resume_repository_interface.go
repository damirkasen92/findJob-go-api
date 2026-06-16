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
		resumeID uint,
	) error

	GetByID(
		ctx context.Context,
		resumeID uint,
	) (*model.Resume, error)

	GetByUserID(
		ctx context.Context,
		userID uint,
	) ([]model.Resume, error)

	GetList(
		ctx context.Context,
	) ([]model.Resume, error)
}
