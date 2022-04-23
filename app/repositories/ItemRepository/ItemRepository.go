// Package ItemRepository.
package ItemRepository

import (
	"TapTake-server/app/models"
	"TapTake-server/app/repositories/RestaurantRepository"
	"TapTake-server/app/services/database"
	"log"
)

// Error classes
type ItemErrorCode int

type ItemError interface {
	Error() string
	Code() ItemErrorCode
}

type iErr struct {
	ErrCode ItemErrorCode
	Err     string
}

func (i iErr) Code() ItemErrorCode {
	return i.ErrCode
}

func (i iErr) Error() string {
	return i.Err
}

const (
	INV_RESTAURANT ItemErrorCode = iota
	INV_PRICE
	INV_QUANTITY
	INV_NAME
	INV_CANCEL
	UNK
)

// Functions

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

func AddNew(item *models.Item) ItemError {

	if !RestaurantRepository.GetById(item.RestaurantId).IsValid() {
		return iErr{ErrCode: INV_RESTAURANT}
	}
	if item.Price < 0 {
		return iErr{ErrCode: INV_PRICE}
	}
	if item.Quantity < 0 {
		return iErr{ErrCode: INV_QUANTITY}
	}
	if item.Name == "" {
		return iErr{ErrCode: INV_NAME}
	}
	if item.CancelGracePeriod < 0 {
		return iErr{ErrCode: INV_CANCEL}
	}

	rows, err := database.Query("INSERT INTO Item (Restaurant, Price, Quantity, Name, Description, CancelGracePeriod) VALUES (?,?,?,?,?,?) RETURNING Id",
		item.RestaurantId, item.Price, item.Quantity, item.Name, item.Description, item.CancelGracePeriod)

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
		err = rows.Scan(&item.Id)

		// Check for errors.
		if err != nil {
			// Notify.
			log.Printf("Couldn't scan Item: %s\n", err.Error())
			return iErr{ErrCode: UNK, Err: err.Error()}
		}
	}
	return nil
}
