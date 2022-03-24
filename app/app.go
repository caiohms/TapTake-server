// Package server.
package app

// Imports.
import (
	"TapTake-server/app/controllers"
	repository "TapTake-server/app/services"
	"fmt"
	"net/http"
	"os"
)

// The Port this Http Server is running on.
var Port = "8090"

// SetupHttpHandlers Setups the Function Handles.
func SetupHttpHandlers() {
	// Notify.
	fmt.Println("Setting up Http Handlers...")

	// Set up endpoints.
	http.HandleFunc("/", controllers.Root)

	// Notify.
	fmt.Println("Http Handlers Set!")
}

// Init the Http Server.
func Init() {

	repository.Init()

	// Notify.
	fmt.Println("Initializing Http Server...")

	// Set up the Http Handlers.
	SetupHttpHandlers()

	// Get the Port from the Environment Variable.
	var EnvPort = os.Getenv("PORT")

	// If we have an Environment Port...
	if len(EnvPort) > 0 {
		// Set the Port.
		Port = EnvPort
	}

	// Notify.
	fmt.Println("Starting Http Server on Port: " + Port)

	// Start a Http server at port.
	var err error = http.ListenAndServe(":"+Port, nil)

	// Check for Error.
	if err != nil {
		// Notify.
		fmt.Println("Error Starting Http Server: " + err.Error())

		// Exit.
		os.Exit(1)
	}

	// Notify.
	fmt.Println("Http Server Started!")
}
