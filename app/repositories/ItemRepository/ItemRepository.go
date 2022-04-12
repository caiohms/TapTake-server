// Package ItemRepository.
package ItemRepository

import (
	"TapTake-server/app/models"
	"TapTake-server/app/repositories/RestaurantRepository"
	"TapTake-server/app/services/database"
	"log"
)

// GetById Gets an Item by Id.
func GetById(Id int) models.Item {
	// Query by Id.
	rows, err := database.Query("SELECT Id, Restaurant, Price, Quantity, Name, Description, CancelGracePeriod FROM Item WHERE id = ?", Id)

	// Check for errors.
	if err != nil {
		// Notify.
		log.Printf("Couldn't query Item by Id %d: %s\n", Id, err.Error())

		// Return empty Item.
		return models.InvalidItem
	}
	defer rows.Close()
	// For each row..
	for rows.Next() {
		// Create a new Item.
		var item models.Item

		// Scan the row.
		err = rows.Scan(&item.Id, &item.RestaurantId, &item.Price, &item.Quantity, &item.Name, &item.Description, &item.CancelGracePeriod)

		// Check for errors.
		if err != nil {
			// Notify.
			log.Printf("Couldn't scan Item by Id %d: %s\n", Id, err.Error())

			// Return empty item.
			return models.InvalidItem
		}

		// Return the item.
		return item
	}

	// In here, there were no results.
	log.Printf("No Item found by Id %d\n", Id)

	// Return empty item.
	return models.InvalidItem
}

// GetRestaurantByitem Gets a Restaurant by an Item.
func GetRestaurantByItem(item models.Item) models.Restaurant {
	// Check for Item.
	if !item.IsValid() {
		// It isn't valid, return invalid restaurant.
		return models.InvalidRestaurant
	}

	// Return Get Restaurant.
	return RestaurantRepository.GetById(item.RestaurantId)
}
