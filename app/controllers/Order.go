package controllers

import (
	"net/http"

	"github.com/go-chi/chi"
)

// Routes for orders
func OrderRouter(r chi.Router) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Seus pedidos"))
	})

	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Pedido Feito"))
	})
}
