package tests

import (
	"TapTake-server/app/repositories/ItemRepository"
	"TapTake-server/app/repositories/RestaurantRepository"
	"TapTake-server/app/repositories/RoleMapRepository"
	"TapTake-server/app/repositories/StatusMapRepository"
	"TapTake-server/app/repositories/UniversityRepository"
	"TapTake-server/app/repositories/UserReferenceRepository"
	"TapTake-server/app/repositories/UserRepository"
	"TapTake-server/app/services/database"
	"fmt"
	"os"
	"testing"
	"time"
)

func TestRepositories(t *testing.T) {

	// Change to root dir
	os.Chdir("..")

	// Create database for this test
	stamp := time.Now().Unix()
	// Use memory mode (does not create local database file)
	database.InitSQLite(fmt.Sprintf("file:db-%d.db?mode=memory", stamp))
	defer database.CloseDB()

	t.Run("Get a valid item", func(t *testing.T) {
		item := ItemRepository.GetById(1)
		if !item.IsValid() {
			t.Fatal("Item is not valid")
		}
	})

	t.Run("Get an invalid item", func(t *testing.T) {
		item := ItemRepository.GetById(3)
		if item.IsValid() {
			t.Fatal("Item is valid")
		}
	})

	t.Run("Get a valid restaurant", func(t *testing.T) {
		rest := RestaurantRepository.GetById(1)
		if !rest.IsValid() {
			t.Fatal("Restaurant is not valid")
		}
	})

	t.Run("Get an invalid restaurant", func(t *testing.T) {
		rest := RestaurantRepository.GetById(2)
		if rest.IsValid() {
			t.Fatal("Restaurant is valid")
		}
	})

	t.Run("Get a valid RoleMap item", func(t *testing.T) {
		obj := RoleMapRepository.GetById(1)
		if !obj.IsValid() {
			t.Fatal("RoleMap is not valid")
		}
	})

	t.Run("Get an invalid RoleMap item", func(t *testing.T) {
		obj := RoleMapRepository.GetById(3)
		if obj.IsValid() {
			t.Fatal("RoleMap is valid")
		}
	})

	t.Run("Get a valid StatusMap item", func(t *testing.T) {
		obj := StatusMapRepository.GetById(1)
		if !obj.IsValid() {
			t.Fatal("RoleMap is not valid")
		}
	})

	t.Run("Get an invalid StatusMap item", func(t *testing.T) {
		obj := StatusMapRepository.GetById(7)
		if obj.IsValid() {
			t.Fatal("RoleMap is valid")
		}
	})

	t.Run("Get a valid University item", func(t *testing.T) {
		obj := UniversityRepository.GetById(1)
		if !obj.IsValid() {
			t.Fatal("University is not valid")
		}
	})

	t.Run("Get an invalid University item", func(t *testing.T) {
		obj := UniversityRepository.GetById(7)
		if obj.IsValid() {
			t.Fatal("University is valid")
		}
	})

	t.Run("Get a valid UserReference item", func(t *testing.T) {
		obj := UserReferenceRepository.GetById(1)
		if !obj.IsValid() {
			t.Fatal("UserReference is not valid")
		}
	})

	t.Run("Get an invalid UserReference item", func(t *testing.T) {
		obj := UserReferenceRepository.GetById(7)
		if obj.IsValid() {
			t.Fatal("UserReference is valid")
		}
	})

	t.Run("Get a valid User item", func(t *testing.T) {
		obj := UserRepository.GetById(1)
		if !obj.IsValid() {
			t.Fatal("User is not valid")
		}
	})

	t.Run("Get an invalid User item", func(t *testing.T) {
		obj := UserRepository.GetById(7)
		if obj.IsValid() {
			t.Fatal("User is valid")
		}
	})
}
