// Package models.
package models

// UserReference struct.
type UserReference struct {
	Id           int
	UserId       int
	UniversityId int
	RestaurantId int
}

// InvalidUserReference is an invalid UserReference.
var InvalidUserReference = UserReference{}

// IsValid check for User Reference.
func (ur UserReference) IsValid() bool {
	return ur.Id > 0 && ur.UserId > 0 && (ur.UniversityId > 0 || ur.RestaurantId > 0)
}
