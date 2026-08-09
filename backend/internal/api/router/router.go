package router

import (
	"net/http"

	"tm/internal/api/handlers"
	"tm/internal/api/middlewares"
)

func NewRouter(h *handlers.Handler) http.Handler {
	mux := http.NewServeMux()

	mux.Handle(
		"POST /new",
		middlewares.Chain(
			http.HandlerFunc(h.CreateHandler),
			middlewares.LogMiddleware,
			middlewares.RecoverMiddleware,
			middlewares.TimeMiddleware,
		),
	)

	return mux
}