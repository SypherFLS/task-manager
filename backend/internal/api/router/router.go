package router

import (
	"net/http"

	"tm/internal/api/handlers"
)

func NewRouter(h *handlers.Handler) http.Handler {
	mux := http.NewServeMux()

	

	return mux
}