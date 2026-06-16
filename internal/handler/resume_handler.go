package handler

import (
	"encoding/json"
	"net/http"

	"github.com/damir/jobfinder/internal/dto"
	"github.com/damir/jobfinder/internal/httpx"
	"github.com/damir/jobfinder/internal/mapper"
	"github.com/damir/jobfinder/internal/service"
	"github.com/damir/jobfinder/internal/validator"
)

type ResumeHandler struct {
	resumeService service.ResumeService
}

func NewResumeHandler(resumeService service.ResumeService) *ResumeHandler {
	return &ResumeHandler{
		resumeService: resumeService,
	}
}

func (h *ResumeHandler) Create(
	w http.ResponseWriter,
	r *http.Request,
) {
	var req dto.CreateResumeRequest

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

	actor := dto.GetActor(r)
	err = h.resumeService.Create(r.Context(), req, *actor)

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
			"message": "resume created",
		},
	)
}

func (h *ResumeHandler) Delete(
	w http.ResponseWriter,
	r *http.Request,
) {
	resumeID, err := httpx.ParseUintParam(r, "resumeID")

	if err != nil {
		httpx.HandleError(
			w,
			err,
		)
		return
	}

	actor := dto.GetActor(r)
	err = h.resumeService.Delete(r.Context(), resumeID, *actor)

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

func (h *ResumeHandler) GetByID(
	w http.ResponseWriter,
	r *http.Request,
) {
	resumeID, err := httpx.ParseUintParam(r, "resumeID")

	if err != nil {
		httpx.HandleError(
			w,
			err,
		)
		return
	}

	resume, err := h.resumeService.GetByID(r.Context(), resumeID)

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
		mapper.ResumeToResponse(*resume),
	)
}

func (h *ResumeHandler) GetList(
	w http.ResponseWriter,
	r *http.Request,
) {
	resumes, err := h.resumeService.GetList(r.Context())

	if err != nil {
		httpx.HandleError(
			w,
			err,
		)

		return
	}

	response := make(
		[]dto.ResumeResponse,
		0,
		len(resumes),
	)

	for _, resume := range resumes {
		response = append(
			response,
			mapper.ResumeToResponse(resume),
		)
	}

	httpx.JSON(
		w,
		http.StatusOK,
		response,
	)
}

func (h *ResumeHandler) MyResumes(
	w http.ResponseWriter,
	r *http.Request,
) {
	actor := dto.GetActor(r)

	resumes, err := h.resumeService.MyResumes(
		r.Context(),
		*actor,
	)

	if err != nil {
		httpx.HandleError(
			w,
			err,
		)

		return
	}

	response := make(
		[]dto.ResumeResponse,
		0,
		len(resumes),
	)

	for _, resume := range resumes {
		response = append(
			response,
			mapper.ResumeToResponse(resume),
		)
	}

	httpx.JSON(
		w,
		http.StatusOK,
		response,
	)
}
