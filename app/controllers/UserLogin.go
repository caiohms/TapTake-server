package controllers

import (
	"net/http"

	"github.com/go-chi/chi"
)

func UserLoginRouter(r chi.Router) {
	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Login logic"))
	})
}
