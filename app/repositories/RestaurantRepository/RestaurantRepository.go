// Package RestaurantRepository.
package RestaurantRepository

import (
	"TapTake-server/app/models"
	"TapTake-server/app/repositories/UniversityRepository"
	"TapTake-server/app/services/database"
	"log"
)

// Error classes
type RestaurantErrorCode int

type RestaurantError interface {
	Error() string
	Code() RestaurantErrorCode
}

type iErr struct {
	ErrCode RestaurantErrorCode
	Err     string
}

func (i iErr) Code() RestaurantErrorCode {
	return i.ErrCode
}

func (i iErr) Error() string {
	return i.Err
}

const (
	INV_NAME RestaurantErrorCode = iota
	INV_UNI
	UNK
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
	defer rows.Close()

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

func AddNew(restaurant *models.Restaurant) RestaurantError {

	if restaurant.Name == "" {
		return iErr{ErrCode: INV_NAME}
	}
	if !UniversityRepository.GetById(restaurant.UniversityId).IsValid() {
		return iErr{ErrCode: INV_UNI}
	}

	rows, err := database.Query("INSERT INTO Restaurant (Name, University) VALUES (?,?) RETURNING Id",
		restaurant.Name, restaurant.UniversityId)

	// Check for errors.
	if err != nil {
		// Notify.
		log.Printf("Couldn't add Item: %s\n", err.Error())
		return iErr{ErrCode: UNK, Err: err.Error()}
	}
	defer rows.Close()
	// For each row..
	for rows.Next() {

		// Scan the row.
		err = rows.Scan(&restaurant.Id)

		// Check for errors.
		if err != nil {
			// Notify.
			log.Printf("Couldn't scan Item: %s\n", err.Error())
			return iErr{ErrCode: UNK, Err: err.Error()}
		}
	}
	return nil
}
