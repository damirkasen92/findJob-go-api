package handler

import (
	"encoding/json"
	"net/http"

	"github.com/damir/jobfinder/internal/dto"
	"github.com/damir/jobfinder/internal/httpx"
	"github.com/damir/jobfinder/internal/middleware"
	"github.com/damir/jobfinder/internal/service"
)

type VacancyHandler struct {
	vacancyService service.VacancyService
}

func NewVacancyHandler(vacancyService service.VacancyService) *VacancyHandler {
	return &VacancyHandler{
		vacancyService: vacancyService,
	}
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

	userID := middleware.GetUserID(r.Context())

	err = h.vacancyService.Create(
		r.Context(),
		req,
		userID,
	)

	if err != nil {
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
