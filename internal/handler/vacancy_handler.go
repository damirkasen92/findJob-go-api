package handler

import (
	"encoding/json"
	"net/http"

	"github.com/damir/jobfinder/internal/dto"
	"github.com/damir/jobfinder/internal/httpx"
	"github.com/damir/jobfinder/internal/mapper"
	"github.com/damir/jobfinder/internal/middleware"
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
		httpx.Error(
			w,
			http.StatusBadRequest,
			err.Error(),
		)
		return
	}

	actor := service.Actor{
		UserID: middleware.GetUserID(r.Context()),
		Role:   middleware.GetRole(r.Context()),
	}
	err = h.vacancyService.Delete(r.Context(), vacancyID, actor)

	if err != nil {
		httpx.Error(
			w,
			http.StatusBadRequest,
			err.Error(),
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
		httpx.Error(
			w,
			http.StatusBadRequest,
			err.Error(),
		)
		return
	}

	vacancy, err := h.vacancyService.GetByID(r.Context(), vacancyID)

	if err != nil {
		httpx.Error(
			w,
			http.StatusBadRequest,
			err.Error(),
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
	vacancies, err := h.vacancyService.List(r.Context())

	if err != nil {
		httpx.Error(
			w,
			http.StatusInternalServerError,
			"internal error",
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
		httpx.Error(
			w,
			http.StatusBadRequest,
			"invalid request",
		)

		return
	}

	err = validator.Validate.Struct(
		req,
	)

	if err != nil {
		httpx.Error(
			w,
			http.StatusBadRequest,
			err.Error(),
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
		if err == service.ErrInvalidSalaryRange {
			httpx.Error(
				w,
				http.StatusBadRequest,
				service.ErrInvalidSalaryRange.Error(),
			)

			return
		}

		httpx.Error(
			w,
			http.StatusInternalServerError,
			"internal error",
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
