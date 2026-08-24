package params 

import (
	"net/http"
	"tm/internal/apperrors"
)

func GetUserID(r *http.Request) (int, error){
	rawUserID := r.Context().Value("userID")
	if rawUserID == "" {
		return 0, apperrors.BlankUserID
	}
	userID := rawUserID.(int)

	return userID, nil
}