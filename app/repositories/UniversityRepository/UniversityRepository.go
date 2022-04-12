// Package UniversityRepository.
package UniversityRepository

import (
	"TapTake-server/app/models"
	"TapTake-server/app/services/database"
	"log"
)

// GetById Gets a University by Id.
func GetById(Id int) models.University {
	// Query by Id.
	rows, err := database.Query("SELECT Id, Name FROM University WHERE id = ?", Id)

	// Check for errors.
	if err != nil {
		// Notify.
		log.Printf("Couldn't query University by Id %d: %s\n", Id, err.Error())

		// Return empty University.
		return models.InvalidUniversity
	}
	defer rows.Close()

	// For each row..
	for rows.Next() {
		// Create a new University.
		var uni models.University

		// Scan the row.
		err = rows.Scan(&uni.Id, &uni.Name)

		// Check for errors.
		if err != nil {
			// Notify.
			log.Printf("Couldn't scan University by Id %d: %s\n", Id, err.Error())

			// Return empty University.
			return models.InvalidUniversity
		}

		// Return the University.
		return uni
	}

	// In here, there were no results.
	log.Printf("No University found by Id %d\n", Id)

	// Return empty university.
	return models.InvalidUniversity
}
