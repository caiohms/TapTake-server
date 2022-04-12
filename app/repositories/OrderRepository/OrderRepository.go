// Package OrderRepository.
package OrderRepository

import (
	"TapTake-server/app/models"
	"TapTake-server/app/repositories/UserRepository"
	"TapTake-server/app/services/database"
	"log"
)

// GetById Gets an Oder by Id.
func GetById(Id int) models.Order {
	// Query by Id.
	rows, err := database.Query("SELECT Id, User, OrderDate, DeliveryDate, CancelDate, FinishDate FROM Order WHERE id = ?", Id)

	// Check for errors.
	if err != nil {
		// Notify.
		log.Printf("Couldn't query Order by Id %d: %s\n", Id, err.Error())

		// Return empty Order.
		return models.InvalidOrder
	}
	defer rows.Close()

	// For each row..
	for rows.Next() {
		// Create a new Item.
		var order models.Order

		// Scan the row.
		err = rows.Scan(&order.Id, &order.UserId, &order.OrderDate, &order.DeliveryDate, &order.CancelDate, &order.FinishDate, &order.StatusId)

		// Check for errors.
		if err != nil {
			// Notify.
			log.Printf("Couldn't scan Order by Id %d: %s\n", Id, err.Error())

			// Return empty order.
			return models.InvalidOrder
		}

		// Return the order.
		return order
	}

	// In here, there were no results.
	log.Printf("No Order found by Id %d\n", Id)

	// Return empty order.
	return models.InvalidOrder
}

// GetUserByOrder Gets a User by an order.
func GetUserByOrder(order models.Order) models.User {
	// Check for Order.
	if !order.IsValid() {
		// Return empty User.
		return models.InvalidUser
	}

	// Get User.
	return UserRepository.GetById(order.UserId)
}
