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
			"ERROR_VALIDATION",
			err.Error(),
		)

	case errors.Is(
		err,
		model.ErrForbidden,
	):
		Error(
			w,
			http.StatusForbidden,
			"ERROR_FORBIDDEN",
			err.Error(),
		)

	case errors.Is(
		err,
		model.ErrNotFound,
	):
		Error(
			w,
			http.StatusNotFound,
			"ERROR_NOT_FOUND",
			err.Error(),
		)

	case errors.Is(
		err,
		model.ErrInvalidCredentials,
	):
		Error(
			w,
			http.StatusUnauthorized,
			"ERROR_INVALID_CREDENTIALS",
			err.Error(),
		)

	case errors.Is(
		err,
		model.ErrUserExists,
	):
		Error(
			w,
			http.StatusBadRequest,
			"ERROR_USER_ALREADY_EXISTS",
			err.Error(),
		)

	case errors.Is(
		err,
		model.ErrInvalidSalaryRange,
	):
		Error(
			w,
			http.StatusBadRequest,
			"ERROR_INVALID_SALARY_RANGE",
			err.Error(),
		)

	case errors.Is(
		err,
		model.ErrInvalidApplicationStatus,
	):
		Error(
			w,
			http.StatusBadRequest,
			"ERROR_INVALID_APPLICATION_STATUS",
			err.Error(),
		)

	default:
		Error(
			w,
			http.StatusInternalServerError,
			"ERROR_INTERNAL_ERROR",
			"internal error",
		)
	}
}
