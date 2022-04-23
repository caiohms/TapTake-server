// Package RoleMapRepository.
package RoleMapRepository

// Imports.
import (
	"TapTake-server/app/models"
	"TapTake-server/app/services/database"
	"log"
)

// GetById Gets a Role Map by Id.
func GetById(Id int) models.RoleMap {
	// Query by Id.
	rows, err := database.Query("SELECT Id, Code, Description FROM RoleMap WHERE id = ?", Id)

	// Check for errors.
	if err != nil {
		// Notify.
		log.Printf("Couldn't query Role Map by Id %d: %s\n", Id, err.Error())

		// Return empty Role Map.
		return models.InvalidRoleMap
	}
	defer rows.Close()

	// For each row..
	for rows.Next() {
		// Create a new Role Map.
		var roleMap models.RoleMap

		// Scan the row.
		err = rows.Scan(&roleMap.Id, &roleMap.Code, &roleMap.Description)

		// Check for errors.
		if err != nil {
			// Notify.
			log.Printf("Couldn't scan Role Map by Id %d: %s\n", Id, err.Error())

			// Return empty Role Map.
			return models.InvalidRoleMap
		}

		// Return the Role Map.
		return roleMap
	}

	// In here, there were no results.
	log.Printf("No Role Map found by Id %d\n", Id)

	// Return empty Role Map.
	return models.InvalidRoleMap
}
