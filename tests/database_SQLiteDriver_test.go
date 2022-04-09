package tests

import (
	"TapTake-server/app/services/database"
	"fmt"
	"log"
	"os"
	"testing"
	"time"
)

func TestDatabaseDriver(t *testing.T) {

	// Change to root dir
	os.Chdir("..")

	log.Println(os.Getwd())

	// Create database for this test
	stamp := time.Now().Unix()
	// Use memory mode (does not create local database file)
	database.InitSQLite(fmt.Sprintf("file:db-%d.db?mode=memory", stamp))
	defer database.CloseDB()

	// Test if receiving right data from default database
	t.Run("Get restaurant list", func(t *testing.T) {

		// Execute Query
		r, err := database.Query("SELECT * FROM Restaurant;")
		defer r.Close()
		if err != nil {
			t.Errorf("Error in database query: %v", err)
		}

		// Read only the first row
		r.Next()
		var id int
		var uni int
		var name string

		// Scan row results to variables
		if err := r.Scan(&id, &uni, &name); err != nil {
			t.Errorf("Error in database scan: %v", err)
		}

		// Error check the variables
		if id != 1 {
			t.Errorf("Exepected ID = 1 got ID = %d", id)
		}

		if uni != 1 {
			t.Errorf("Exepected University = 1 got University = %d", uni)
		}

		if name != "Lanchonete DCE" {
			t.Errorf("Exepected Name = \"Lanchonete DCE\" got Name = \"%s\"", name)
		}
	})
}
