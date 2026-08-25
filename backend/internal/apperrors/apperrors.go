package apperrors

import "errors"

var (
	ErrNotFound        = errors.New("task not found")
	ErrNothingToUpdate = errors.New("nothing to update")
	ErrConflict        = errors.New("conflict")
	BadQuery           = errors.New("bad query")
	BlankUserID        = errors.New("blank user id")
	WrongUserID        = errors.New("Wrong user ID")
	WrongSignMethod    = errors.New("unexpected signing method")
	InvalidToken       = errors.New("invalid token")
)
