// Define Main Package.
package main

// Imports.
import (
	"TapTake-server/app"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

var port int = 8090

// Main Function.
func main() {
	// Get the Port.
	EnvPort := os.Getenv("PORT")

	if len(EnvPort) > 0 {
		var err error
		port, err = strconv.Atoi(EnvPort)
		if err != nil {
			log.Fatal(err)
		}
	}
	log.Printf("Starting server on port %d", port)
	err := http.ListenAndServe(":"+fmt.Sprint(port), app.Init())

	if err != nil {
		log.Fatal(err)
	}
}
