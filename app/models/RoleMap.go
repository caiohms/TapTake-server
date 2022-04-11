// Package models.
package models

// Role Enum.
type Role int

// Create Role Values.
const (
	RoleStudent Role = iota
	RoleRestaurant
	RoleInvalid // Used for Invalid Calls.
)

// RoleMap This Struct is used for keeping Role Data.
type RoleMap struct {
	Id          int
	Code        Role
	Description string
}

// InvalidRoleMap This Struct is used for keeping Invalid Role Data.
var InvalidRoleMap = RoleMap{
	Id:          0,
	Code:        RoleInvalid,
	Description: "Invalid Role",
}

// IsValid check for RoleMap.
func (role RoleMap) IsValid() bool {
	return role.Id > 0 && role.Code != RoleInvalid
}
