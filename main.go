// Define Main Package.
package main

// Imports.
import (
	"TapTake-server/database"
	"TapTake-server/server"
)

// Main Function.
func main() {
	// First, init the database.
	database.Init()

	// Then, init the http server.
	server.Init()
}
