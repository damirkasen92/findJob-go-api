package dto

type CreateVacancyRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`

	SalaryFrom int `json:"salary_from"`
	SalaryTo   int `json:"salary_to"`
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
