package api

import (
	"net/http"
	"time"
)

func WithRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.Header.Get("X-Request-ID")
		if id == "" {
			id = time.Now().UTC().Format("20060102150405.000000000")
		}
		w.Header().Set("X-Request-ID", id)
		next.ServeHTTP(w, r)
	})
}
func WithRecovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if recover() != nil {
				http.Error(w, "internal error", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
func ChainMiddleware(h http.Handler) http.Handler { return WithRequestID(WithRecovery(h)) }
