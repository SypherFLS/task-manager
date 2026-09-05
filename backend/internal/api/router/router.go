package router

import (
	"net/http"

	"tm/internal/api/handlers"
	"tm/internal/api/middlewares"
	"tm/internal/auth"
	"tm/internal/config"
)

func NewRouter(h *handlers.Handler, jwtManager *auth.JWTManager, cfg config.Config) http.Handler {
	root := http.NewServeMux()

	public := http.NewServeMux()
	private := http.NewServeMux()

	public.Handle("POST /register", http.HandlerFunc(h.UserRegisterHandler))
	public.Handle("POST /login", http.HandlerFunc(h.UserLoginHandler))

	private.Handle("POST /tasks", http.HandlerFunc(h.TaskCreateHandler))
	private.Handle("GET /tasks", http.HandlerFunc(h.TaskReadHandler))
	private.Handle("DELETE /task/{id}", http.HandlerFunc(h.TaskDeleteHandler))
	private.Handle("PATCH /task/{id}", http.HandlerFunc(h.TaskUpdateHandler))

	publicChain := middlewares.CommonChain(
		public,
		cfg.Server.Timeout,
	)

	privateChain := middlewares.CommonChain(
		middlewares.AuthMiddleware(jwtManager)(
			private,
		),
		cfg.Server.Timeout,
	)

	root.Handle("/api/", http.StripPrefix("/api", privateChain))
	root.Handle("/auth/", http.StripPrefix("/auth", publicChain))

	return root
}
