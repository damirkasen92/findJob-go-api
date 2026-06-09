package httpx

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type Response struct {
	Data  any    `json:"data,omitempty"`
	Error string `json:"error,omitempty"`
}

func ParseUintParam(
	r *http.Request,
	name string,
) (uint, error) {
	idStr := chi.URLParam(
		r,
		name,
	)

	id64, err := strconv.ParseUint(
		idStr,
		10,
		64,
	)

	if err != nil {
		return 0, err
	}

	return uint(id64), nil
}

func JSON(
	w http.ResponseWriter,
	status int,
	data any,
) {

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	w.WriteHeader(status)

	_ = json.NewEncoder(w).Encode(
		Response{
			Data: data,
		},
	)
}

func Error(
	w http.ResponseWriter,
	status int,
	message string,
) {

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	w.WriteHeader(status)

	_ = json.NewEncoder(w).Encode(
		Response{
			Error: message,
		},
	)
}
