package repository

import (
	"context"

	"github.com/damir/jobfinder/internal/dto"
	"github.com/damir/jobfinder/internal/model"
	"gorm.io/gorm"
)

type resumeRepository struct {
	db *gorm.DB
}

func NewResumeRepository(db *gorm.DB) ResumeRepository {
	return &resumeRepository{
		db: db,
	}
}

func (r *resumeRepository) Create(
	ctx context.Context,
	resume *model.Resume,
) error {
	return r.db.
		WithContext(ctx).
		Create(resume).
		Error
}

func (r *resumeRepository) Update(
	ctx context.Context,
	dto dto.UpdateResumeRequest,
) error {
	return r.db.
		WithContext(ctx).
		Model(&model.Resume{}).
		Where("id = ?", dto.Id).
		Updates(map[string]interface{}{
			"title":  dto.Title,
			"about":  dto.About,
			"skills": dto.Skills,
		}).
		Error
}

func (r *resumeRepository) Delete(
	ctx context.Context,
	resumeID uint,
) error {
	return r.db.
		WithContext(ctx).
		Delete(
			&model.Resume{},
			resumeID,
		).
		Error
}

func (r *resumeRepository) GetByID(
	ctx context.Context,
	resumeID uint,
) (*model.Resume, error) {
	var resume model.Resume

	err := r.db.
		WithContext(ctx).
		First(&resume, resumeID).
		Error

	if err != nil {
		return nil, err
	}

	return &resume, nil
}

func (r *resumeRepository) GetList(
	ctx context.Context,
) ([]model.Resume, error) {
	var resumes []model.Resume

	err := r.db.
		WithContext(ctx).
		Model(&model.Resume{}).
		Find(&resumes).
		Error

	if err != nil {
		return nil, err
	}

	return resumes, nil
}

func (r *resumeRepository) GetByUserID(
	ctx context.Context,
	userID uint,
) ([]model.Resume, error) {
	var resumes []model.Resume

	db := r.db.
		WithContext(ctx).
		Model(
			&model.Resume{},
		).
		Where("user_id", userID)

	err := db.Find(
		&resumes,
	).Error

	return resumes, err
}
