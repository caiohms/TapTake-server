// Package models.
package models

// University struct.
type University struct {
	Id   int
	Name string
}

// InvalidUniversity is an invalid university.
var InvalidUniversity = University{}

// IsValid check for University.
func (u University) IsValid() bool {
	return u.Id > 0
}
