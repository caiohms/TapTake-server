// Package models.
package models

// Order struct.
type Order struct {
	Id           int
	UserId       int
	OrderDate    int64
	DeliveryDate int64
	CancelDate   int64
	FinishDate   int64
	StatusId     int
}

// InvalidOrder is an invalid Order.
var InvalidOrder = Order{}

// IsValid check for Order.
func (o Order) IsValid() bool {
	return o.Id > 0
}
