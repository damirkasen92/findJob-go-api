package mapper

import (
	"encoding/json"

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

func UserToResponse(u model.User) dto.UserResponse {
	return dto.UserResponse{
		ID:    u.ID,
		Email: u.Email,
	}
}

func ResumeToResponse(r model.Resume) dto.ResumeResponse {
	var skills []string
	_ = json.Unmarshal(r.Skills, &skills)

	return dto.ResumeResponse{
		ID:     r.ID,
		Title:  r.Title,
		About:  r.About,
		Skills: skills,
		User:   UserToResponse(r.User),
	}
}

func ApplicationToResponse(a model.Application) dto.ApplicationResponse {
	return dto.ApplicationResponse{
		ID:        a.ID,
		VacancyID: a.VacancyID,
		Status:    a.Status,
		CreatedAt: a.CreatedAt,
		Resume:    ResumeToResponse(a.Resume),
	}
}
