package repository

import (
	"context"

	"github.com/damir/jobfinder/internal/model"
	"gorm.io/gorm"
)

type vacancyRepository struct {
	db *gorm.DB
}

func NewVacancyRepository(db *gorm.DB) VacancyRepository {
	return &vacancyRepository{
		db: db,
	}
}

func (r *vacancyRepository) Create(
	ctx context.Context,
	vacancy *model.Vacancy,
) error {
	return r.db.
		WithContext(ctx).
		Create(vacancy).
		Error
}

func (r *vacancyRepository) Update(
	ctx context.Context,
	vacancy *model.Vacancy,
) error {
	return r.db.
		WithContext(ctx).
		Save(vacancy).
		Error
}

func (r *vacancyRepository) Delete(
	ctx context.Context,
	id uint,
) error {
	return r.db.
		WithContext(ctx).
		Delete(
			&model.Vacancy{},
			id,
		).
		Error
}

func (r *vacancyRepository) GetByID(
	ctx context.Context,
	id uint,
) (*model.Vacancy, error) {
	var vacancy model.Vacancy

	err := r.db.
		WithContext(ctx).
		First(&vacancy, id).
		Error

	if err != nil {
		return nil, err
	}

	return &vacancy, nil
}

func (r *vacancyRepository) List(
	ctx context.Context,
) ([]model.Vacancy, error) {
	var vacancies []model.Vacancy

	err := r.db.
		WithContext(ctx).
		Find(&vacancies).
		Error

	return vacancies, err
}
