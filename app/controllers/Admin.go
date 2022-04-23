package controllers

import (
	"net/http"

	"github.com/go-chi/chi"
)

func AdminRouter(r chi.Router) {

	// use Auth
	// r.Use()
	r.Route("/university", func(r chi.Router) {

		// Add new restaurant
		r.Post("/restaurant", func(w http.ResponseWriter, r *http.Request) {

		})
	})

	r.Route("/restaurant", func(r chi.Router) {
		r.Route("/menu", func(r chi.Router) {

			// Get all items on the menu
			r.Get("/", func(w http.ResponseWriter, r *http.Request) {

			})

			// Add new item on menu
			r.Post("/", func(w http.ResponseWriter, r *http.Request) {

			})

			// Specific item operations
			r.Route("/{item}", func(r chi.Router) {

				// Get specific Item
				r.Get("/", func(w http.ResponseWriter, r *http.Request) {

				})

				// Update specific Item
				r.Put("/", func(w http.ResponseWriter, r *http.Request) {

				})

				// Delete specific Item
				r.Delete("/", func(w http.ResponseWriter, r *http.Request) {

				})
			})
		})
	})

}
