package controllers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
)

// Router for restaurant controller
func RestaurantRouter(r chi.Router) {

	// Open endpoints
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Restaurantes"))
	})

	// Get specific restaurant data
	r.Get("/{restaurante}", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf("Voce é %s", chi.URLParam(r, "restaurante"))))
	})
	protected := r.Group(nil)
	protected.Use(middleware.Logger)
	protected.Post("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Cadastrar Restaurante"))
	})
	protected.Put("/{restaurante}", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf("Update %s", chi.URLParam(r, "restaurante"))))
	})

	// Delete restaurant
	protected.Delete("/{restaurante}", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf("Delete %s", chi.URLParam(r, "restaurante"))))
	})

}
