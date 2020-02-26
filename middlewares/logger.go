package middlewares

import (
	"log"
	"net/http"
)

// Logging use for log a http request
func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%v] Called with method %v \n", r.URL.Path, r.Method)
		next.ServeHTTP(w, r)
	})
}
