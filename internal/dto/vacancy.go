package dto

type CreateVacancyRequest struct {
	Title       string `json:"title" validate:"required,min=3,max=100"`
	Description string `json:"description" validate:"required,min=10,max=2000"`

	SalaryFrom int `json:"salary_from" validate:"required,gt=0"`
	SalaryTo   int `json:"salary_to" validate:"required,gt=0"`
}

type UpdateVacancyRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`

	SalaryFrom int `json:"salary_from"`
	SalaryTo   int `json:"salary_to"`
}

type VacancyResponse struct {
	ID uint `json:"id"`

	Title       string `json:"title"`
	Description string `json:"description"`

	SalaryFrom int `json:"salary_from"`
	SalaryTo   int `json:"salary_to"`

	CreatedBy uint `json:"created_by"`
}
