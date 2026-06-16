package handler

import (
	"encoding/json"
	"net/http"

	"github.com/damir/jobfinder/internal/dto"
	"github.com/damir/jobfinder/internal/httpx"
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
