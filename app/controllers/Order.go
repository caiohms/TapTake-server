package controllers

import (
	"net/http"

	"github.com/go-chi/chi"
)

// Routes for orders
func OrderRouter(r chi.Router) {

	// Use auth for all the endpoints
	//r.Use()

	// Group all user routes
	r.Route("/user", func(r chi.Router) {
		// Get all orders
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Seus pedidos"))
		})

		// Get specific order
		r.Get("/{pedido}", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Seus pedidos"))
		})

		// Create a new order
		r.Post("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Pedido Feito"))
		})

		// Update specific order
		r.Put("/{pedido}", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Update"))
		})
	})

	// Group all restaurant routes
	r.Route("/restaurant", func(r chi.Router) {

		// Get all orders for this restaurant
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Seus pedidos"))
		})

		// Get specific order
		r.Get("/{pedido}", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Seus pedidos"))
		})

		// Update specific order
		r.Put("/{pedido}", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Update"))
		})
	})

}
