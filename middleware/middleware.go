package middleware

import (
	"log"
	"net/http"
)

func Middleware(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("requested URL : %s", r.URL.RequestURI())
		h.ServeHTTP(w, r)
	}
}
