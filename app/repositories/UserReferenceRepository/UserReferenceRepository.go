// Package UserReferenceRepository.
package UserReferenceRepository

import (
	"TapTake-server/app/models"
	"TapTake-server/app/repositories/RestaurantRepository"
	"TapTake-server/app/repositories/UniversityRepository"
	"TapTake-server/app/repositories/UserRepository"
	"TapTake-server/app/services/database"
	"database/sql"
	"fmt"
	"log"
)

// Error classes
type UserReferenceErrorCode int

type UserReferenceError interface {
	Error() string
	Code() UserReferenceErrorCode
}

type iErr struct {
	ErrCode UserReferenceErrorCode
	Err     string
}

func (i iErr) Code() UserReferenceErrorCode {
	return i.ErrCode
}

func (i iErr) Error() string {
	return i.Err
}

const (
	INV_CONFIG UserReferenceErrorCode = iota
	INV_USR
	UNK
)

// GetById Gets a User Reference by Id.
func GetById(Id int) models.UserReference {

	var rows *sql.Rows
	var err error

	if database.DBType == database.SQLite3 {
		// Query by Id.
		rows, err = database.Query("SELECT Id, Ifnull(University, 0) ,Ifnull(Restaurant, 0), User FROM UserReference WHERE id = ?", Id)
	}

	// Check for errors.
	if err != nil {
		// Notify.
		log.Printf("Couldn't query User Reference by Id %d: %s\n", Id, err.Error())

		// Return empty User Reference.
		return models.InvalidUserReference
	}
	defer rows.Close()

	// For each row..
	for rows.Next() {
		// Create a new User Reference.
		var userref models.UserReference

		// Scan the row.
		err = rows.Scan(&userref.Id, &userref.UniversityId, &userref.RestaurantId, &userref.UserId)

		// Check for errors.
		if err != nil {
			// Notify.
			log.Printf("Couldn't scan User Reference by Id %d: %s\n", Id, err.Error())

			// Return empty User Reference.
			return models.InvalidUserReference
		}

		// Return the User Reference.
		return userref
	}

	// In here, there were no results.
	log.Printf("No User Reference found by Id %d\n", Id)

	// Return empty User Reference.
	return models.InvalidUserReference
}

// GetUserByUserReference Gets a User by a User Reference.
func GetUserByUserReference(userref models.UserReference) models.User {
	// Check for User Reference.
	if !userref.IsValid() {
		// Return empty User.
		return models.InvalidUser
	}

	// Get User.
	return UserRepository.GetById(userref.UserId)
}

// GetUniversityByUserReference Gets a University by a User Reference.
func GetUniversityByUserReference(userref models.UserReference) models.University {
	// Check for User Reference.
	if !userref.IsValid() || userref.UniversityId < 1 || UserRepository.GetUserRole(GetUserByUserReference(userref)) != models.RoleStudent {
		// It isn't valid, return invalid university.
		return models.InvalidUniversity
	}

	// Return Get University.
	return UniversityRepository.GetById(userref.UniversityId)
}

// GetRestaurantByUserReference Gets a Restaurant by a User Reference.
func GetRestaurantByUserReference(userref models.UserReference) models.Restaurant {
	// Check for User Reference.
	if !userref.IsValid() || userref.RestaurantId < 1 || UserRepository.GetUserRole(GetUserByUserReference(userref)) != models.RoleRestaurant {
		// It isn't valid, return invalid restaurant.
		return models.InvalidRestaurant
	}

	// Return Get Restaurant.
	return RestaurantRepository.GetById(userref.RestaurantId)
}

func AddNew(ur *models.UserReference) UserReferenceError {

	if !UserRepository.GetById(ur.UserId).IsValid() {
		return iErr{ErrCode: INV_USR}
	}

	if ur.RestaurantId > 0 && ur.UniversityId > 0 {
		return iErr{ErrCode: INV_CONFIG}
	}

	if ur.RestaurantId == 0 && ur.UniversityId == 0 {
		return iErr{ErrCode: INV_CONFIG}
	}

	Uni := "NULL"
	Restaurant := "NULL"

	if ur.UniversityId > 0 {
		Uni = fmt.Sprint(ur.UniversityId)
	}
	if ur.RestaurantId > 0 {
		Restaurant = fmt.Sprint(ur.RestaurantId)
	}

	rows, err := database.Query("INSERT INTO UserReference (University, Restaurant, User) VALUES (?,?,?) RETURNING Id",
		Uni, Restaurant, ur.UserId)

	// Check for errors.
	if err != nil {
		// Notify.
		log.Printf("Couldn't add Item: %s\n", err.Error())
		return iErr{ErrCode: UNK, Err: err.Error()}
	}
	defer rows.Close()
	// For each row..
	for rows.Next() {

		// Scan the row.
		err = rows.Scan(&ur.Id)

		// Check for errors.
		if err != nil {
			// Notify.
			log.Printf("Couldn't scan Item: %s\n", err.Error())
			return iErr{ErrCode: UNK, Err: err.Error()}
		}
	}

	return nil
}
