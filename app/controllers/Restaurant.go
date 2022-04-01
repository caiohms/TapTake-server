package controllers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

// Router for restaurant controller
func RestaurantRouter(r chi.Router) {

	// Get a list of all restaurants
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Restaurantes"))
	})

	// Get data for one restaurant
	r.Get("/{restaurante}", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf("Voce é %s", chi.URLParam(r, "restaurante"))))
	})

	// Group all endpoints that require authentication
	protected := r.Group(nil)
	// Add authentication middleware
	//protected.Use()

	// Create a new restaurant
	protected.Post("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Cadastrar Restaurante"))
	})

	// Update a restaurant
	protected.Put("/{restaurante}", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf("Update %s", chi.URLParam(r, "restaurante"))))
	})

	// Delete a restaurant
	protected.Delete("/{restaurante}", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf("Delete %s", chi.URLParam(r, "restaurante"))))
	})

}
