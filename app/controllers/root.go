// Package endpoints.
package controllers

// Imports.
import (
	"TapTake-server/app/models"
	database "TapTake-server/app/services"
	"TapTake-server/app/utils"
	"fmt"
	"net/http"
)

// Root Gets the Root '/' of the Http Server.
func Root(resp http.ResponseWriter, request *http.Request) {
	// Check HTTP Method.
	if !utils.CheckRequestMethod(request, http.MethodGet, resp) {
		// Method is incorrect, return.
		return
	}

	// The Result String.
	var Rst models.Test = models.Test{
		Id: 0,
		Values: []string{
			"Hello! This is the TapTake Server.",
		},
	}

	// Run Simple Query.
	sqlRows, err := database.Query("SELECT 1")

	// Check for Error.
	if err != nil {
		// Notify.
		Rst.Values = append(Rst.Values, fmt.Sprintf("Error during query: %s", err.Error()))

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
			Rst.Values = append(Rst.Values, fmt.Sprintf("Error during scan: %s", err.Error()))

			// Continue to next row.
			continue
		}

		// Print it!
		Rst.Values = append(Rst.Values, fmt.Sprintf("Result: %d", Result))
	}

	// Write Response.
	utils.WriteResponse(resp, http.StatusOK, Rst)
}
