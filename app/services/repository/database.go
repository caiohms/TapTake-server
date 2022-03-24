// Package database.
package repository

// Imports.
import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

// Database connection.
var db *sql.DB

// ConnectionString for the local deployment.
var ConnectionString = "host=localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable"

// Init Inits the Postgres database.
func Init() {
	// Notify.
	fmt.Println("Initializing database...")

	// Grab the database connection string from the environment.
	var EnvConnectionString = os.Getenv("DATABASE_URL")

	// Check for EnvConnectionString.
	if len(EnvConnectionString) > 0 {
		// Use the environment variable.
		ConnectionString = EnvConnectionString
	}

	// Defines the Error, if any.
	var err error

	// Opens the Database Connection using postgres
	db, err = sql.Open("postgres", ConnectionString)

	// Check for error.
	if err != nil {
		// Print the Error.
		fmt.Println("Error Connecting to Postgres: " + err.Error())

		// Exit the program.
		os.Exit(1)
	}

	// Notify.
	fmt.Println("Database initialized.")
}

// Query executes a query.
func Query(query string, args ...interface{}) (*sql.Rows, error) {
	// Execute the query.
	return db.Query(query, args...)
}

/*
 * Only export "top level" functions
 * use SQL only in this file
 */
func GetRestaurante(nome string) {
	Query("SELECT * FROM restaurantes WHERE name=$1", nome)
}
