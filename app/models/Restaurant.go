// Package models.
package models

// Restaurant struct.
type Restaurant struct {
	Id           int
	Name         string
	UniversityId int
}

// InvalidRestaurant is an invalid restaurant.
var InvalidRestaurant = Restaurant{}

// IsValid check for Restaurant.
func (r Restaurant) IsValid() bool {
	return r.Id > 0
}
