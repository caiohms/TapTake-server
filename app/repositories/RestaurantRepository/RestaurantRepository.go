// Package RestaurantRepository.
package RestaurantRepository

import (
	"TapTake-server/app/models"
	"TapTake-server/app/repositories/UniversityRepository"
	"TapTake-server/app/services/database"
	"log"
)

// GetById Gets a Restaurant by Id.
func GetById(Id int) models.Restaurant {
	// Query by Id.
	rows, err := database.Query("SELECT Id, University, Name FROM Restaurant WHERE id = ?", Id)

	// Check for errors.
	if err != nil {
		// Notify.
		log.Printf("Couldn't query Restaurant by Id %d: %s\n", Id, err.Error())

		// Return empty Restaurant.
		return models.InvalidRestaurant
	}

	// For each row..
	for rows.Next() {
		// Create a new Restaurant.
		var restaurant models.Restaurant

		// Scan the row.
		err = rows.Scan(&restaurant.Id, &restaurant.UniversityId, &restaurant.Name)

		// Check for errors.
		if err != nil {
			// Notify.
			log.Printf("Couldn't scan Restaurant by Id %d: %s\n", Id, err.Error())

			// Return empty Restaurant.
			return models.InvalidRestaurant
		}

		// Return the Restaurant.
		return restaurant
	}

	// In here, there were no results.
	log.Printf("No Restaurant found by Id %d\n", Id)

	// Return empty restaurant.
	return models.InvalidRestaurant
}

// GetUniversityByRestaurant Gets a University by Restaurant.
func GetUniversityByRestaurant(restaurant models.Restaurant) models.University {
	// Check for Restaurant.
	if !restaurant.IsValid() {
		// It isn't valid, return invalid university.
		return models.InvalidUniversity
	}

	// Return Get by Id.
	return UniversityRepository.GetById(restaurant.UniversityId)
}