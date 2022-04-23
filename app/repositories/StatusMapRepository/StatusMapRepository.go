// Package StatusMapRepository.
package StatusMapRepository

// Imports.
import (
	"TapTake-server/app/models"
	"TapTake-server/app/services/database"
	"log"
)

// GetById Gets a Status Map by Id.
func GetById(Id int) models.StatusMap {
	// Query by Id.
	rows, err := database.Query("SELECT Id, Code, Description FROM StatusMap WHERE id = ?", Id)

	// Check for errors.
	if err != nil {
		// Notify.
		log.Printf("Couldn't query Status Map by Id %d: %s\n", Id, err.Error())

		// Return empty Status Map.
		return models.InvalidStatusMap
	}
	defer rows.Close()

	// For each row..
	for rows.Next() {
		// Create a new Status Map.
		var statusMap models.StatusMap

		// Scan the row.
		err = rows.Scan(&statusMap.Id, &statusMap.Code, &statusMap.Description)

		// Check for errors.
		if err != nil {
			// Notify.
			log.Printf("Couldn't scan Status Map by Id %d: %s\n", Id, err.Error())

			// Return empty Status Map.
			return models.InvalidStatusMap
		}

		// Return the Status Map.
		return statusMap
	}

	// In here, there were no results.
	log.Printf("No Status Map found by Id %d\n", Id)

	// Return empty Status Map.
	return models.InvalidStatusMap
}
