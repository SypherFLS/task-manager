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
		"GET /users",
		middlewares.CommonChain(
			http.HandlerFunc(h.ReadHandler),
		),
	)
	mux.Handle(
		"DELETE /user/{id}",
		middlewares.CommonChain(
			http.HandlerFunc(h.DeleteHandler),
		),
	)
	mux.Handle(
		"PATCH /user/{id}",
		middlewares.CommonChain(
			http.HandlerFunc(h.UpdateHandler),
		),
	)

	return mux
}