package dto

type CreateApplicationRequest struct {
	ResumeID  uint `json:"resume_id"`
	VacancyID uint `json:"vacancy_id"`
}
