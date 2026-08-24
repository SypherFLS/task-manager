package params

import (
	"net/http"
	"strconv"
	"tm/internal/apperrors"
)

type FilPage struct {
	Priority *string

	Limit  int
	Offset int
}

func ParseQuery(r *http.Request) (FilPage, error) {
	var fp FilPage
	priority := r.URL.Query().Get("priority")

	switch priority {
	case "high", "middle", "low":
		fp.Priority = &priority
	case "":
	default:
		return fp, apperrors.BadQuery
	}

	limit := r.URL.Query().Get("limit")
	if limit != "" {
		limInt, err := strconv.Atoi(limit)
		if err != nil || limInt <= 0 || limInt > 1000 {
			return fp, apperrors.BadQuery
		}
		fp.Limit = limInt
	} else {
		fp.Limit = 10
	}

	offset := r.URL.Query().Get("offset")
	if offset != "" {
		offsetInt, err := strconv.Atoi(offset)
		if err != nil || offsetInt < 0 {
			return fp, apperrors.BadQuery
		}
		fp.Offset = offsetInt
	} else {
		fp.Offset = 0
	}

	return fp, nil
}
