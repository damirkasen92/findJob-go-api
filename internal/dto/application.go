package dto

import (
	"time"

	"github.com/damir/jobfinder/internal/model"
)

type CreateApplicationRequest struct {
	ResumeID  uint `json:"resume_id"`
	VacancyID uint `json:"vacancy_id"`
}

type UpdateApplicationStatusRequest struct {
	Status model.ApplicationStatus `json:"status"`
}

type ApplicationResponse struct {
	ID        uint                    `json:"id"`
	VacancyID uint                    `json:"vacancy_id"`
	Status    model.ApplicationStatus `json:"status"`
	CreatedAt time.Time               `json:"created_at"`
	Resume    ResumeResponse          `json:"resume"`
}
