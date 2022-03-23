// Package endpoints.
package endpoints

// Imports.
import (
	"TapTake-server/database"
	"fmt"
	"net/http"
)

// Root Gets the Root '/' of the Http Server.
func Root(resp http.ResponseWriter, request *http.Request) {
	// Prints to WebBrowser the Response.
	fmt.Fprintf(resp, "Hello!\n")

	// Run Simple Query.
	sqlRows, err := database.Query("SELECT 1")

	// Check for Error.
	if err != nil {
		// Notify.
		fmt.Fprintf(resp, "Error during query: %s\n", err.Error())

		// Stop running function.
		return
	}

	// For each row...
	for sqlRows.Next() {
		// Create Result.
		var Result = 0

		// Scan the row.
		err = sqlRows.Scan(&Result)

		// Check for Error.
		if err != nil {
			// Notify.
			fmt.Fprintf(resp, "Error during scan: %s\n", err.Error())

			// Continue to next row.
			continue
		}

		// Print it!
		fmt.Fprintf(resp, "Row: %d\n", Result)
	}
}
