package query

type VacancySort string

const (
	SortNewest     VacancySort = "newest"
	SortSalaryAsc  VacancySort = "salary_asc"
	SortSalaryDesc VacancySort = "salary_desc"
)

func ParseVacancySort(value string) VacancySort {
	switch VacancySort(value) {
	case SortNewest, SortSalaryAsc, SortSalaryDesc:
		return VacancySort(value)
	default:
		return SortNewest
	}
}

func GetSortingForDB(value VacancySort) string {
	switch value {
	case SortSalaryAsc:
		return "salary_from asc"
	case SortSalaryDesc:
		return "salary_to desc"
	default:
		return "created_at desc"
	}
}
