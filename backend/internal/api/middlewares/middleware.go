package middlewares

import (
	"net/http"
	"tm/internal/api/utils/selfwriter"
	"log"
	"time"
)

type Middleware func(http.Handler) http.Handler 

func Chain(h http.Handler, m ...Middleware) http.Handler {
	for i := len(m); i >= 0; i-- {
		h = m[i](h)
	}

	return h
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