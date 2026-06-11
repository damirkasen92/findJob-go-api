package repository

import (
	"context"

	"github.com/damir/jobfinder/internal/model"
	"github.com/damir/jobfinder/internal/query"
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
	filter query.VacancyFilter,
) ([]model.Vacancy, error) {
	var vacancies []model.Vacancy

	db := r.db.
		WithContext(ctx).
		Model(
			&model.Vacancy{},
		)

	if filter.Search != "" {
		db = db.Where(
			"title ILIKE ?",
			"%"+filter.Search+"%",
		)
	}

	if filter.SalaryFrom > 0 {
		db = db.Where(
			"salary_from >= ?",
			filter.SalaryFrom,
		)
	}

	if filter.SalaryTo > 0 {
		db = db.Where(
			"salary_to <= ?",
			filter.SalaryTo,
		)
	}

	if filter.CreatedBy > 0 {
		db = db.Where(
			"created_by = ?",
			filter.CreatedBy,
		)
	}

	db = db.Order(
		query.GetSortingForDB(filter.Sort),
	)

	offset :=
		(filter.Page - 1) *
			filter.Limit

	db = db.
		Limit(filter.Limit).
		Offset(offset)

	err := db.Find(
		&vacancies,
	).Error

	return vacancies, err
}
