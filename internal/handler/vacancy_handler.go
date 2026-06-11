package handler

import (
	"encoding/json"
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
		httpx.HandleError(
			w,
			err,
		)
		return
	}

	actor := service.Actor{
		UserID: middleware.GetUserID(r.Context()),
		Role:   middleware.GetRole(r.Context()),
	}
	err = h.vacancyService.Delete(r.Context(), vacancyID, actor)

	if err != nil {
		httpx.HandleError(
			w,
			err,
		)
		return
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
		httpx.HandleError(
			w,
			err,
		)
		return
	}

	vacancy, err := h.vacancyService.GetByID(r.Context(), vacancyID)

	if err != nil {
		httpx.HandleError(
			w,
			err,
		)
		return
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
	vacancies, err := h.vacancyService.List(r.Context(), filter)

	if err != nil {
		httpx.HandleError(
			w,
			err,
		)

		return
	}

	response := make(
		[]dto.VacancyResponse,
		0,
		len(vacancies),
	)

	for _, vacancy := range vacancies {
		response = append(
			response,
			mapper.VacancyToResponse(vacancy),
		)
	}

	httpx.JSON(
		w,
		http.StatusOK,
		response,
	)
}

func (h *VacancyHandler) Create(
	w http.ResponseWriter,
	r *http.Request,
) {
	var req dto.CreateVacancyRequest

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		httpx.HandleError(
			w,
			err,
		)

		return
	}

	err = validator.ValidateStruct(
		req,
	)

	if err != nil {
		httpx.HandleError(
			w,
			err,
		)

		return
	}

	userID := middleware.GetUserID(r.Context())

	err = h.vacancyService.Create(
		r.Context(),
		req,
		userID,
	)

	if err != nil {
		httpx.HandleError(
			w,
			err,
		)

		return
	}

	httpx.JSON(
		w,
		http.StatusCreated,
		map[string]string{
			"message": "vacancy created",
		},
	)
}
