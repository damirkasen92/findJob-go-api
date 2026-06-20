package handler

import (
	"encoding/json"
	"net/http"

	"github.com/damir/jobfinder/internal/dto"
	"github.com/damir/jobfinder/internal/httpx"
	"github.com/damir/jobfinder/internal/mapper"
	"github.com/damir/jobfinder/internal/service"
)

type ApplicationHandler struct {
	applicationService service.ApplicationService
}

func NewApplicationHandler(applicationService service.ApplicationService) *ApplicationHandler {
	return &ApplicationHandler{
		applicationService: applicationService,
	}
}

func (h *ApplicationHandler) Create(
	w http.ResponseWriter,
	r *http.Request,
) {
	var req dto.CreateApplicationRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		panic(err)
	}

	actor := dto.GetActor(r)

	if err := h.applicationService.Create(
		r.Context(),
		req,
		*actor,
	); err != nil {
		panic(err)
	}

	httpx.JSON(
		w,
		http.StatusCreated,
		map[string]string{
			"message": "application created",
		},
	)
}

func (h *ApplicationHandler) ListByVacancy(
	w http.ResponseWriter,
	r *http.Request,
) {
	vacancyID, err := httpx.ParseUintParam(r, "vacancyID")

	if err != nil {
		panic(err)
	}

	actor := dto.GetActor(r)

	applications, err := h.applicationService.ListByVacancy(r.Context(), vacancyID, *actor)

	if err != nil {
		panic(err)
	}

	responses := make([]dto.ApplicationResponse, 0, len(applications))
	for _, app := range applications {
		responses = append(responses, mapper.ApplicationToResponse(app))
	}

	httpx.JSON(
		w,
		http.StatusOK,
		responses,
	)
}

func (h *ApplicationHandler) UpdateStatus(
	w http.ResponseWriter,
	r *http.Request,
) {
	var dto dto.UpdateApplicationStatusRequest

	appId, err := httpx.ParseUintParam(r, "id")

	if err != nil {
		panic(err)
	}

	if err = json.NewDecoder(r.Body).Decode(&dto); err != nil {
		panic(err)
	}

	if err = h.applicationService.UpdateStatus(r.Context(), appId, dto); err != nil {
		panic(err)
	}

	httpx.JSON(
		w,
		http.StatusOK,
		map[string]string{
			"message": "status updated",
		},
	)
}

func (h *ApplicationHandler) GetByUser(
	w http.ResponseWriter,
	r *http.Request,
) {
	actor := dto.GetActor(r)
	applications, err := h.applicationService.ListByUser(r.Context(), *actor)

	if err != nil {
		panic(err)
	}

	responses := make([]dto.ApplicationResponse, 0, len(applications))
	for _, app := range applications {
		responses = append(responses, mapper.ApplicationToResponse(app))
	}

	httpx.JSON(
		w,
		http.StatusOK,
		responses,
	)
}
