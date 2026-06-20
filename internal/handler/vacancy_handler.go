package handler

import (
	"encoding/json"
	"math"
	"net/http"

	"github.com/damir/jobfinder/internal/dto"
	"github.com/damir/jobfinder/internal/httpx"
	"github.com/damir/jobfinder/internal/mapper"
	"github.com/damir/jobfinder/internal/middleware"
	"github.com/damir/jobfinder/internal/query"
	"github.com/damir/jobfinder/internal/service"
	"github.com/damir/jobfinder/internal/validator"
)

type VacancyHandler struct {
	vacancyService service.VacancyService
}

func NewVacancyHandler(vacancyService service.VacancyService) *VacancyHandler {
	return &VacancyHandler{
		vacancyService: vacancyService,
	}
}

func (h *VacancyHandler) Delete(
	w http.ResponseWriter,
	r *http.Request,
) {
	vacancyID, err := httpx.ParseUintParam(r, "vacancyID")

	if err != nil {
		panic(err)
	}

	actor := dto.GetActor(r)

	if err = h.vacancyService.Delete(r.Context(), vacancyID, *actor); err != nil {
		panic(err)
	}

	httpx.JSON(
		w,
		http.StatusNoContent,
		"",
	)
}

func (h *VacancyHandler) GetByID(
	w http.ResponseWriter,
	r *http.Request,
) {
	vacancyID, err := httpx.ParseUintParam(r, "vacancyID")

	if err != nil {
		panic(err)
	}

	vacancy, err := h.vacancyService.GetByID(r.Context(), vacancyID)

	if err != nil {
		panic(err)
	}

	httpx.JSON(
		w,
		http.StatusOK,
		mapper.VacancyToResponse(*vacancy),
	)
}

func (h *VacancyHandler) GetList(
	w http.ResponseWriter,
	r *http.Request,
) {
	filter := query.ParseVacancyFilter(r)
	vacancies, total, err := h.vacancyService.List(r.Context(), filter)

	if err != nil {
		panic(err)
	}

	vacancyResponse := make(
		[]dto.VacancyResponse,
		0,
		len(vacancies),
	)

	for _, vacancy := range vacancies {
		vacancyResponse = append(
			vacancyResponse,
			mapper.VacancyToResponse(vacancy),
		)
	}

	pages := int(
		math.Ceil(
			float64(total) /
				float64(filter.Limit),
		),
	)

	paginatedResponse :=
		httpx.PaginatedResponse[dto.VacancyResponse]{
			Data: vacancyResponse,

			Meta: httpx.Meta{
				Page:  filter.Page,
				Limit: filter.Limit,

				Total: total,
				Pages: pages,
			},
		}

	httpx.JSON(
		w,
		http.StatusOK,
		paginatedResponse,
	)
}

func (h *VacancyHandler) Create(
	w http.ResponseWriter,
	r *http.Request,
) {
	var req dto.CreateVacancyRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		panic(err)
	}

	if err := validator.ValidateStruct(
		req,
	); err != nil {
		panic(err)
	}

	userID := middleware.GetUserID(r.Context())

	if err := h.vacancyService.Create(
		r.Context(),
		req,
		userID,
	); err != nil {
		panic(err)
	}

	httpx.JSON(
		w,
		http.StatusCreated,
		map[string]string{
			"message": "vacancy created",
		},
	)
}
