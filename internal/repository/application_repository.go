package repository

import (
	"context"

	"github.com/damir/jobfinder/internal/model"
	"gorm.io/gorm"
)

type applicationRepository struct {
	db *gorm.DB
}

func NewApplicationRepository(db *gorm.DB) ApplicationRepository {
	return &applicationRepository{
		db: db,
	}
}

func (r *applicationRepository) Create(
	ctx context.Context,
	application *model.Application,
) error {
	return r.db.
		WithContext(ctx).
		Create(application).
		Error
}

func (r *applicationRepository) Update(
	ctx context.Context,
	application *model.Application,
) error {
	return r.db.
		WithContext(ctx).
		Save(application).
		Error
}

func (r *applicationRepository) GetByID(
	ctx context.Context,
	applicationID uint,
) (*model.Application, error) {
	var application model.Application

	err := r.db.
		WithContext(ctx).
		First(&application, applicationID).
		Error

	return &application, err
}

func (r *applicationRepository) Exists(
	ctx context.Context,
	resumeID uint,
	vacancyID uint,
) (bool, error) {
	var count int64

	err := r.db.
		WithContext(ctx).
		Model(&model.Application{}).
		Where(
			"resume_id = ? AND vacancy_id = ?",
			resumeID,
			vacancyID,
		).
		Count(&count).
		Error

	return count > 0, err
}

func (r *applicationRepository) ListByUser(
	ctx context.Context,
	userID uint,
) ([]model.Application, error) {
	var applications []model.Application

	db := r.db.
		WithContext(ctx).
		Model(&model.Application{}).
		Preload("Resume").
		Preload("Resume.User").
		Where("user_id", userID)

	err := db.Find(&applications).
		Error

	return applications, err
}

func (r *applicationRepository) ListByVacancy(
	ctx context.Context,
	vacancyID uint,
) ([]model.Application, error) {
	var applications []model.Application

	db := r.db.
		WithContext(ctx).
		Model(&model.Application{}).
		Where("vacancy_id", vacancyID).
		Preload("Resume").
		Preload("Resume.User")

	err := db.Find(&applications).
		Error

	return applications, err
}
