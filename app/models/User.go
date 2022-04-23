// Package models.
package models

// User struct.
type User struct {
	Id       int
	Name     string
	Email    string
	Password string
	RoleId   int
}

// InvalidUser checks if a user is invalid.
var InvalidUser = User{}

// IsValid check for User.
func (u User) IsValid() bool {
	return u.Id > 0
}
