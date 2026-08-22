package errormapping

import (
	"errors"
	"net/http"
	"tm/internal/apperrors"
)

func StatusFromError(err error) int {
	switch {
	case errors.Is(err, apperrors.ErrNotFound):
		return http.StatusNotFound

	case errors.Is(err, apperrors.ErrNothingToUpdate):
		return http.StatusBadRequest

	case errors.Is(err, apperrors.ErrConflict):
		return http.StatusConflict

	default:
		return http.StatusInternalServerError
	}
}
