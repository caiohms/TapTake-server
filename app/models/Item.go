// Package models.
package models

// Item struct.
type Item struct {
	Id                int
	Name              string
	Description       string
	Price             float64
	Quantity          int
	CancelGracePeriod int
	RestaurantId      int
}

// InvalidItem is an invalid item.
var InvalidItem = Item{}

// IsValid check for Item.
func (i Item) IsValid() bool {
	return i.Id > 0 && i.RestaurantId > 0
}
