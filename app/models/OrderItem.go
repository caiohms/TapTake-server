// Package models.
package models

// OrderItem struct.
type OrderItem struct {
	Id       int
	OrderId  int
	ItemId   int
	Quantity int
	Price    float64
}

// InvalidOrderItem is an invalid order item.
var InvalidOrderItem = OrderItem{}

// IsValid check for order item.
func (o OrderItem) IsValid() bool {
	return o.Id > 0 && o.OrderId > 0
}
