package errormapping

import (
	"net/http"
	"tm/internal/apperrors"
	"errors"
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