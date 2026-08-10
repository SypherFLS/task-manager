package middlewares

import (
	"context"
	"log"
	"net/http"
	"time"
	"tm/internal/api/utils/helpers"
	"tm/internal/api/utils/selfwriter"
)

type Middleware func(http.Handler) http.Handler 

func Chain(h http.Handler, m ...Middleware) http.Handler {
	for i := len(m); i >= 0; i-- {
		h = m[i](h)
	}

	return h
}

func CommonChain(h http.Handler) http.Handler {
	return Chain(
		h,
		LogMiddleware,
		RecoverMiddleware,
		TimeMiddleware,
	)
}

func TimeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
		defer cancel()

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}


func LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		current := time.Now()
		sw := &selfwriter.SelfWriter {
			ResponseWriter : w,
			Code : 200,
		}
		log.Printf("handler %v started \n", r.URL.String()) // сразу весь запрос с query
		next.ServeHTTP(sw, r)
		log.Printf("handler %v finished with code %v and time %v \n", r.URL.Path, sw.Code, time.Since(current))
	})
}

func RecoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		defer func() {
			if err := recover(); err != nil {
				helpers.WriteError(w, 500, "panic")
				return
			}
		}()
		next.ServeHTTP(w, r)
	})
} 

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//TODO AUTH

		next.ServeHTTP(w, r)
	})
}