package apperrors

import "errors"

var (
    ErrNotFound        = errors.New("task not found")
    ErrNothingToUpdate = errors.New("nothing to update")
    ErrConflict        = errors.New("conflict")
)