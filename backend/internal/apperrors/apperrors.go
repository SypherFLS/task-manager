package apperrors

import "errors"

var (
	ErrNotFound        = errors.New("task not found")
	ErrNothingToUpdate = errors.New("nothing to update")
	ErrConflict        = errors.New("conflict")
	BadQuery           = errors.New("bad query")
	BlankUserID        = errors.New("blank user id")
)
