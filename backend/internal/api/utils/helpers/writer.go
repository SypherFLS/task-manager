package helpers

import (
	"encoding/json"
	"log"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func WriteJSON(w http.ResponseWriter, code int, data any) {
	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	w.WriteHeader(code)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("failed encoding with error: %v", err)
	}
}

func WriteError(w http.ResponseWriter, code int, msg string) {
	erro := ErrorResponse{
		Error: msg,
	}
	WriteJSON(w, code, erro)
}
