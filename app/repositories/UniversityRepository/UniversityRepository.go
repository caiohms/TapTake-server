// Package UniversityRepository.
package UniversityRepository

import (
	"TapTake-server/app/models"
	"TapTake-server/app/services/database"
	"log"
)

// Error classes
type UniversityErrorCode int

type UniversityError interface {
	Error() string
	Code() UniversityErrorCode
}

type iErr struct {
	ErrCode UniversityErrorCode
	Err     string
}

func (i iErr) Code() UniversityErrorCode {
	return i.ErrCode
}

func (i iErr) Error() string {
	return i.Err
}

const (
	INV_NAME UniversityErrorCode = iota
	UNK
)

// GetById Gets a University by Id.
func GetById(Id int) models.University {
	// Query by Id.
	rows, err := database.Query("SELECT Id, Name FROM University WHERE id = ?", Id)

	// Check for errors.
	if err != nil {
		// Notify.
		log.Printf("Couldn't query University by Id %d: %s\n", Id, err.Error())

		// Return empty University.
		return models.InvalidUniversity
	}
	defer rows.Close()

	// For each row..
	for rows.Next() {
		// Create a new University.
		var uni models.University

		// Scan the row.
		err = rows.Scan(&uni.Id, &uni.Name)

		// Check for errors.
		if err != nil {
			// Notify.
			log.Printf("Couldn't scan University by Id %d: %s\n", Id, err.Error())

			// Return empty University.
			return models.InvalidUniversity
		}

		// Return the University.
		return uni
	}

	// In here, there were no results.
	log.Printf("No University found by Id %d\n", Id)

	// Return empty university.
	return models.InvalidUniversity
}

func AddNew(uni *models.University) UniversityError {

	if uni.Name == "" {
		return iErr{ErrCode: INV_NAME}
	}
	rows, err := database.Query("INSERT INTO University (Name) VALUES (?) RETURNING Id",
		uni.Name)

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
		err = rows.Scan(&uni.Id)

		// Check for errors.
		if err != nil {
			// Notify.
			log.Printf("Couldn't scan Item: %s\n", err.Error())
			return iErr{ErrCode: UNK, Err: err.Error()}
		}
	}
	return nil
}
