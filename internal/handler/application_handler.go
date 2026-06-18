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

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		httpx.HandleError(
			w,
			err,
		)

		return
	}

	actor := dto.GetActor(r)
	err = h.applicationService.Create(r.Context(), req, *actor)

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
		httpx.HandleError(
			w,
			err,
		)

		return
	}

	actor := dto.GetActor(r)
	applications, err := h.applicationService.ListByVacancy(r.Context(), vacancyID, *actor)

	if err != nil {
		httpx.HandleError(
			w,
			err,
		)

		return
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
		httpx.HandleError(
			w,
			err,
		)

		return
	}

	err = json.NewDecoder(r.Body).Decode(&dto)

	if err != nil {
		httpx.HandleError(
			w,
			err,
		)

		return
	}

	err = h.applicationService.UpdateStatus(r.Context(), appId, dto)

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
		httpx.HandleError(
			w,
			err,
		)

		return
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
