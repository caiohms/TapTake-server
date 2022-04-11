// Package UserReferenceRepository.
package UserReferenceRepository

import (
	"TapTake-server/app/models"
	"TapTake-server/app/repositories/RestaurantRepository"
	"TapTake-server/app/repositories/UniversityRepository"
	"TapTake-server/app/repositories/UserRepository"
	"TapTake-server/app/services/database"
	"log"
)

// GetById Gets a User Reference by Id.
func GetById(Id int) models.UserReference {
	// Query by Id.
	rows, err := database.Query("SELECT Id, University, Restaurant, User FROM UserReference WHERE id = ?", Id)

	// Check for errors.
	if err != nil {
		// Notify.
		log.Printf("Couldn't query User Reference by Id %d: %s\n", Id, err.Error())

		// Return empty User Reference.
		return models.InvalidUserReference
	}

	// For each row..
	for rows.Next() {
		// Create a new User Reference.
		var userref models.UserReference

		// Scan the row.
		err = rows.Scan(&userref.Id, &userref.UniversityId, &userref.RestaurantId, &userref.UserId)

		// Check for errors.
		if err != nil {
			// Notify.
			log.Printf("Couldn't scan User Reference by Id %d: %s\n", Id, err.Error())

			// Return empty User Reference.
			return models.InvalidUserReference
		}

		// Return the User Reference.
		return userref
	}

	// In here, there were no results.
	log.Printf("No User Reference found by Id %d\n", Id)

	// Return empty User Reference.
	return models.InvalidUserReference
}

// GetUserByUserReference Gets a User by a User Reference.
func GetUserByUserReference(userref models.UserReference) models.User {
	// Check for User Reference.
	if !userref.IsValid() {
		// Return empty User.
		return models.InvalidUser
	}

	// Get User.
	return UserRepository.GetById(userref.UserId)
}

// GetUniversityByUserReference Gets a University by a User Reference.
func GetUniversityByUserReference(userref models.UserReference) models.University {
	// Check for User Reference.
	if !userref.IsValid() || userref.UniversityId < 1 || UserRepository.GetUserRole(GetUserByUserReference(userref)) != models.RoleStudent {
		// It isn't valid, return invalid university.
		return models.InvalidUniversity
	}

	// Return Get University.
	return UniversityRepository.GetById(userref.UniversityId)
}

// GetRestaurantByUserReference Gets a Restaurant by a User Reference.
func GetRestaurantByUserReference(userref models.UserReference) models.Restaurant {
	// Check for User Reference.
	if !userref.IsValid() || userref.RestaurantId < 1 || UserRepository.GetUserRole(GetUserByUserReference(userref)) != models.RoleRestaurant {
		// It isn't valid, return invalid restaurant.
		return models.InvalidRestaurant
	}

	// Return Get Restaurant.
	return RestaurantRepository.GetById(userref.RestaurantId)
}
