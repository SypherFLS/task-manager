package middlewares

import (
	"context"
	"log"
	"net/http"
	"runtime/debug"
	"strings"
	"time"
	"tm/internal/api/utils/helpers"
	"tm/internal/api/utils/params"
	"tm/internal/api/utils/selfwriter"
	"tm/internal/auth"

	_ "github.com/golang-jwt/jwt/v5"
)

type Middleware func(http.Handler) http.Handler

func Chain(h http.Handler, m ...Middleware) http.Handler {
	for i := len(m) - 1; i >= 0; i-- {
		h = m[i](h)
	}

	return h
}

func CommonChain(h http.Handler) http.Handler {
	return Chain(
		h,
		LogMiddleware,
		RecoverMiddleware,
		TimeoutMiddleware,
	)
}

func TimeoutMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
		defer cancel()

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		current := time.Now()
		sw := &selfwriter.SelfWriter{
			ResponseWriter: w,
			Code:           200,
		}
		log.Printf("handler %v started \n", r.URL.String())
		next.ServeHTTP(sw, r)
		log.Printf("handler %v finished with code %v and time %v \n", r.URL.Path, sw.Code, time.Since(current))
	})
}

func RecoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("PANIC: %v\n", err)
				debug.PrintStack()

				helpers.WriteError(w, 500, "panic")
				return
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func TraceMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//TODO REQUEST ID TRACE

		next.ServeHTTP(w, r)
	})
}

const bearer = "Bearer "

func AuthMiddleware(jwtManager *auth.JWTManager) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")

			if authHeader == "" {
				helpers.WriteError(w, http.StatusUnauthorized, "invalid authorization header")
				return
			}
			if !strings.HasPrefix(authHeader, bearer) {
				helpers.WriteError(w, http.StatusUnauthorized, "invalid authorization header")
				return
			}

			token := strings.TrimPrefix(authHeader, bearer)

			userID, err := jwtManager.Validate(token)
			if err != nil {
				helpers.WriteError(w, http.StatusUnauthorized, "invalid authorization header")
				return
			}
			ctx := context.WithValue(r.Context(), params.UserIDKey, userID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
