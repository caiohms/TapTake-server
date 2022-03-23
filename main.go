// Define Main Package.
package main

// Imports.
import (
	"fmt"
	"net/http"
)

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

	// Start a Http server at port 8090.
	http.ListenAndServe(":8090", nil)
}
