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
		middlewares.CommonChain(
			http.HandlerFunc(h.CreateHandler),
		),
	)
	mux.Handle(
		"GET /tasks",
		middlewares.CommonChain(
			http.HandlerFunc(h.ReadHandler),
		),
	)
	mux.Handle(
		"DELETE /task/{id}",
		middlewares.CommonChain(
			http.HandlerFunc(h.DeleteHandler),
		),
	)
	mux.Handle(
		"PATCH /task/{id}",
		middlewares.CommonChain(
			http.HandlerFunc(h.UpdateHandler),
		),
	)

	return mux
}