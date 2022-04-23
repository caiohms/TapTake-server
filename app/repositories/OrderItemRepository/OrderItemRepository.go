// Package OrderItemRepository.
package OrderItemRepository

import (
	"TapTake-server/app/models"
	"TapTake-server/app/repositories/ItemRepository"
	"TapTake-server/app/repositories/OrderRepository"
	"TapTake-server/app/services/database"
	"log"
)

// GetById Gets an Oder by Id.
func GetById(Id int) models.OrderItem {
	// Query by Id.
	rows, err := database.Query("SELECT Id, Order, Item, Quantity, Price FROM OrderItem WHERE id = ?", Id)

	// Check for errors.
	if err != nil {
		// Notify.
		log.Printf("Couldn't query Order Item by Id %d: %s\n", Id, err.Error())

		// Return empty Order item.
		return models.InvalidOrderItem
	}
	defer rows.Close()

	// For each row..
	for rows.Next() {
		// Create a new Item.
		var orderitem models.OrderItem

		// Scan the row.
		err = rows.Scan(&orderitem.Id, &orderitem.OrderId, &orderitem.ItemId, &orderitem.Quantity, &orderitem.Price)

		// Check for errors.
		if err != nil {
			// Notify.
			log.Printf("Couldn't scan Order Item by Id %d: %s\n", Id, err.Error())

			// Return empty order.
			return models.InvalidOrderItem
		}

		// Return the order.
		return orderitem
	}

	// In here, there were no results.
	log.Printf("No Order Item found by Id %d\n", Id)

	// Return empty order.
	return models.InvalidOrderItem
}

// GetOrderByOrderItem Gets an Order by an Item.
func GetOrderByOrderItem(orderitem models.OrderItem) models.Order {
	// Check for Order Item.
	if !orderitem.IsValid() {
		// Invalid order!
		return models.InvalidOrder
	}

	// Get by Id.
	return OrderRepository.GetById(orderitem.OrderId)
}

// GetItemByOrderItem Gets an Item by an Order.
func GetItemByOrderItem(orderitem models.OrderItem) models.Item {
	// Check for Order Item.
	if !orderitem.IsValid() {
		// Invalid order!
		return models.InvalidItem
	}

	// Get by Id.
	return ItemRepository.GetById(orderitem.ItemId)
}
