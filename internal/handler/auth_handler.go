package handler

import (
	"encoding/json"
	"net/http"

	"github.com/damir/jobfinder/internal/dto"
	"github.com/damir/jobfinder/internal/httpx"
	"github.com/damir/jobfinder/internal/middleware"
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
	if err := json.NewDecoder(
		r.Body,
	).Decode(&req); err != nil {
		panic(err)
	}

	// do registering business logic
	if err := h.userService.Register(
		r.Context(),
		req,
	); err != nil {
		panic(err)
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

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		panic(err)
	}

	accessToken, refreshToken, err := h.userService.Login(
		r.Context(),
		req,
	)

	if err != nil {
		panic(err)
	}

	httpx.JSON(
		w,
		http.StatusOK,
		map[string]string{
			"access_token":  accessToken,
			"refresh_token": refreshToken,
		},
	)
}

func (h *AuthHandler) Refresh(
	w http.ResponseWriter,
	r *http.Request,
) {
	var req dto.RefreshRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		panic(err)
	}

	accessToken, err := h.userService.Refresh(r.Context(), req.RefreshToken)

	if err != nil {
		panic(err)
	}

	httpx.JSON(
		w,
		http.StatusOK,
		map[string]string{
			"access_token": accessToken,
		},
	)
}

func (h *AuthHandler) Me(
	w http.ResponseWriter,
	r *http.Request,
) {
	// GetUserID - helper to retrive data from ctx by key
	userID := middleware.GetUserID(
		r.Context(),
	)

	// get a user from db (layers, layers and more layers)
	user, err := h.userService.GetByID(
		r.Context(),
		userID,
	)

	if err != nil {
		panic(err)
	}

	// send a simple json
	httpx.JSON(
		w,
		http.StatusOK,
		map[string]any{
			"id":    user.ID,
			"email": user.Email,
			"role":  user.Role,
		},
	)
}
