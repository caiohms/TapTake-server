// Define Main Package.
package main

// Imports.
import (
	"fmt"
	"net/http"
	"os"
)

// The Port this Http Server is running on.
var Port = "8090"

// GetRoot Gets the Root '/' of the Http Server.
func GetRoot(resp http.ResponseWriter, request *http.Request) {
	// Prints to WebBrowser the Response.
	fmt.Fprintf(resp, "Hello!")
}

// Main Function.
func main() {
	// Just print to Console.
	fmt.Println("Hello World!")

	// Specify that root '/' calls GetRoot function.
	http.HandleFunc("/", GetRoot)

	// Get the Port from the Environment Variable.
	var EnvPort = os.Getenv("PORT")

	// If we have an Environment Port...
	if len(EnvPort) > 0 {
		// Set the Port.
		Port = EnvPort
	}

	// Start a Http server at port.
	http.ListenAndServe(":"+Port, nil)
}
