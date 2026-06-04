package handler

import (
	"encoding/json"
	"net/http"

	"github.com/damir/jobfinder/internal/dto"
	"github.com/damir/jobfinder/internal/httpx"
	"github.com/damir/jobfinder/internal/service"
)

type AuthHandler struct {
	userService service.UserService
}

func NewAuthHandler(
	userService service.UserService,
) *AuthHandler {

	return &AuthHandler{
		userService: userService,
	}
}

func (h *AuthHandler) Register(
	w http.ResponseWriter,
	r *http.Request,
) {

	var req dto.RegisterRequest

	// recieve our data
	err := json.NewDecoder(
		r.Body,
	).Decode(&req)

	if err != nil {
		httpx.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	// do registering business logic
	err = h.userService.Register(
		r.Context(),
		req,
	)

	// raw errors
	if err != nil {

		if err == service.ErrUserExists {
			httpx.Error(w, http.StatusConflict, err.Error())
			return
		}

		httpx.Error(w, http.StatusInternalServerError, err.Error())

		return
	}

	// send a response with status code 201
	httpx.JSON(
		w,
		http.StatusCreated,
		map[string]string{
			"message": "user created",
		},
	)
}

func (h *AuthHandler) Login(
	w http.ResponseWriter,
	r *http.Request,
) {
	var req dto.LoginRequest

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		httpx.Error(
			w,
			http.StatusBadRequest,
			"invalid request",
		)

		return
	}

	token, err := h.userService.Login(
		r.Context(),
		req,
	)

	if err != nil {
		if err == service.ErrInvalidCredentials {
			httpx.Error(
				w,
				http.StatusUnauthorized,
				err.Error(),
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
		http.StatusOK,
		map[string]string{
			"access_token": token,
		},
	)
}
