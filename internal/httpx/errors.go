package httpx

import (
	"errors"
	"net/http"

	"github.com/damir/jobfinder/internal/model"
)

func HandleError(
	w http.ResponseWriter,
	err error,
) {
	switch {
	case errors.Is(
		err,
		model.ErrValidation,
	):
		Error(
			w,
			http.StatusBadRequest,
			err.Error(),
		)

	case errors.Is(
		err,
		model.ErrForbidden,
	):
		Error(
			w,
			http.StatusForbidden,
			err.Error(),
		)

	case errors.Is(
		err,
		model.ErrNotFound,
	):
		Error(
			w,
			http.StatusNotFound,
			err.Error(),
		)

	case errors.Is(
		err,
		model.ErrInvalidCredentials,
	):
		Error(
			w,
			http.StatusUnauthorized,
			err.Error(),
		)

	case errors.Is(
		err,
		model.ErrUserExists,
	):
		Error(
			w,
			http.StatusBadRequest,
			err.Error(),
		)

	case errors.Is(
		err,
		model.ErrInvalidSalaryRange,
	):
		Error(
			w,
			http.StatusBadRequest,
			err.Error(),
		)

	default:
		Error(
			w,
			http.StatusInternalServerError,
			"internal error",
		)
	}
}
