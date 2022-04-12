// Package UserRepository.
package UserRepository

import (
	"TapTake-server/app/models"
	"TapTake-server/app/repositories/RoleMapRepository"
	"TapTake-server/app/services/database"
	"log"
)

// GetById Gets a User by Id.
func GetById(Id int) models.User {
	// Query by Id.
	rows, err := database.Query("SELECT Id, Email, Password, Role, Name FROM Users WHERE id = ?", Id)

	// Check for errors.
	if err != nil {
		// Notify.
		log.Printf("Couldn't query Users by Id %d: %s\n", Id, err.Error())

		// Return empty Users.
		return models.InvalidUser
	}
	defer rows.Close()

	// For each row..
	for rows.Next() {
		// Create a new User.
		var user models.User

		// Scan the row.
		err = rows.Scan(&user.Id, &user.Email, &user.Password, &user.RoleId, &user.Name)

		// Check for errors.
		if err != nil {
			// Notify.
			log.Printf("Couldn't scan Users by Id %d: %s\n", Id, err.Error())

			// Return empty User.
			return models.InvalidUser
		}

		// Return the User.
		return user
	}

	// In here, there were no results.
	log.Printf("No User found by Id %d\n", Id)

	// Return empty User.
	return models.InvalidUser
}

// GetUserRole gets the Role of a User.
func GetUserRole(user models.User) models.Role {
	// If the user isn't valid...
	if !user.IsValid() {
		// Return Role Invalid.
		return models.RoleInvalid
	}

	// Get the Role, return the code.
	return RoleMapRepository.GetById(user.RoleId).Code
}
