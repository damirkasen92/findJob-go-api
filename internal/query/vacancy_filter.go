package query

import (
	"net/http"
	"strconv"
)

type VacancyFilter struct {
	Page  int
	Limit int

	Search string

	SalaryFrom int
	SalaryTo   int

	CreatedBy uint

	Sort VacancySort
}

func ParseVacancyFilter(
	r *http.Request,
) VacancyFilter {

	page, _ := strconv.Atoi(
		r.URL.Query().Get("page"),
	)

	if page <= 0 {
		page = 1
	}

	limit, _ := strconv.Atoi(
		r.URL.Query().Get("limit"),
	)

	if limit <= 0 {
		limit = 20
	}

	salaryFrom, _ := strconv.Atoi(
		r.URL.Query().Get("salary_from"),
	)

	salaryTo, _ := strconv.Atoi(
		r.URL.Query().Get("salary_to"),
	)

	createdBy64, _ := strconv.ParseUint(
		r.URL.Query().Get("created_by"),
		10,
		64,
	)

	return VacancyFilter{
		Page:  page,
		Limit: limit,

		Search: r.URL.Query().Get(
			"search",
		),

		SalaryFrom: salaryFrom,
		SalaryTo:   salaryTo,

		CreatedBy: uint(createdBy64),

		Sort: ParseVacancySort(r.URL.Query().Get(
			"sort",
		)),
	}
}
