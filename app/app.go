// Package server.
package app

// Imports.
import (
	"TapTake-server/app/controllers"
	"TapTake-server/app/utils/middlewares"

	"github.com/go-chi/chi"
)

// Init the Http Server.
func Init() chi.Router {

	// Register handlers
	router := chi.NewRouter()

	router.Use(middlewares.HttpLogger)
	router.Route("/restaurant", controllers.RestaurantRouter)
	router.Route("/order", controllers.OrderRouter)

	return router
}
