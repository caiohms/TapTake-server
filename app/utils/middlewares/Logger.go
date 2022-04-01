package middlewares

import (
	"log"
	"net/http"
)

func HttpLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		log.Printf("%s %s\n", req.Method, req.URL.Path)
		next.ServeHTTP(w, req)
	})
}
