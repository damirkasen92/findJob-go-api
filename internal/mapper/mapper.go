package mapper

import (
	"github.com/damir/jobfinder/internal/dto"
	"github.com/damir/jobfinder/internal/model"
)

func VacancyToResponse(
	v model.Vacancy,
) dto.VacancyResponse {
	return dto.VacancyResponse{
		ID: v.ID,

		Title:       v.Title,
		Description: v.Description,

		SalaryFrom: v.SalaryFrom,
		SalaryTo:   v.SalaryTo,

		CreatedBy: v.CreatedBy,
	}
}
