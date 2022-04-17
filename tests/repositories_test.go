package tests

import (
	"TapTake-server/app/models"
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

func TestStaticRepositories(t *testing.T) {

	// Change to root dir
	os.Chdir("..")

	// Create database for this test
	stamp := time.Now().Unix()
	// Use memory mode (does not create local database file)
	database.InitSQLite(fmt.Sprintf("file:db-%d.db?mode=memory", stamp))
	defer database.CloseDB()

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

func TestItemRepository(t *testing.T) {

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

	t.Run("Add new Item", func(t *testing.T) {
		item := models.Item{
			RestaurantId:      1,
			Price:             15,
			Quantity:          30,
			Name:              "Hamburguer",
			Description:       "Dois hamburgueres alface queijo molho especial...",
			CancelGracePeriod: 10,
		}
		err := ItemRepository.AddNew(&item)

		if err != nil {
			t.Log(err)
		}

		if !item.IsValid() {
			t.Fatalf("Added Item is invalid")
		}

	})

	t.Run("Add new Item with invalid restaurant", func(t *testing.T) {
		item := models.Item{
			RestaurantId:      0,
			Price:             15,
			Quantity:          30,
			Name:              "Hamburguer",
			Description:       "Dois hamburgueres alface queijo molho especial...",
			CancelGracePeriod: 10,
		}
		err := ItemRepository.AddNew(&item)
		if err.Code() != ItemRepository.INV_RESTAURANT {
			t.Fatalf("Error is nil")
		}
		if item.IsValid() {
			t.Fatalf("Added Item is valid")
		}

	})

	t.Run("Add new Item with invalid price", func(t *testing.T) {
		item := models.Item{
			RestaurantId:      1,
			Price:             -1,
			Quantity:          30,
			Name:              "Hamburguer",
			Description:       "Dois hamburgueres alface queijo molho especial...",
			CancelGracePeriod: 10,
		}
		err := ItemRepository.AddNew(&item)
		if err.Code() != ItemRepository.INV_PRICE {
			t.Fatalf("Error is nil")
		}
		if item.IsValid() {
			t.Fatalf("Added Item is valid")
		}

	})

	t.Run("Add new Item with invalid quantity", func(t *testing.T) {
		item := models.Item{
			RestaurantId:      1,
			Price:             1,
			Quantity:          -1,
			Name:              "Hamburguer",
			Description:       "Dois hamburgueres alface queijo molho especial...",
			CancelGracePeriod: 10,
		}
		err := ItemRepository.AddNew(&item)
		if err.Code() != ItemRepository.INV_QUANTITY {
			t.Fatalf("Error is nil")
		}
		if item.IsValid() {
			t.Fatalf("Added Item is valid")
		}

	})

	t.Run("Add new Item with invalid name", func(t *testing.T) {
		item := models.Item{
			RestaurantId:      1,
			Price:             1,
			Quantity:          1,
			Name:              "",
			Description:       "Dois hamburgueres alface queijo molho especial...",
			CancelGracePeriod: 10,
		}
		err := ItemRepository.AddNew(&item)
		if err.Code() != ItemRepository.INV_NAME {
			t.Fatalf("Error is nil")
		}
		if item.IsValid() {
			t.Fatalf("Added Item is valid")
		}

	})

	t.Run("Add new Item with invalid cancel", func(t *testing.T) {
		item := models.Item{
			RestaurantId:      1,
			Price:             1,
			Quantity:          1,
			Name:              "A",
			Description:       "Dois hamburgueres alface queijo molho especial...",
			CancelGracePeriod: -1,
		}
		err := ItemRepository.AddNew(&item)
		if err.Code() != ItemRepository.INV_CANCEL {
			t.Fatalf("Error is nil")
		}
		if item.IsValid() {
			t.Fatalf("Added Item is valid")
		}

	})
}

func TestRestaurantRepository(t *testing.T) {
	// Change to root dir
	os.Chdir("..")

	// Create database for this test
	stamp := time.Now().Unix()
	// Use memory mode (does not create local database file)
	database.InitSQLite(fmt.Sprintf("file:db-%d.db?mode=memory", stamp))
	defer database.CloseDB()

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

	t.Run("Add new restaurant", func(t *testing.T) {
		rest := models.Restaurant{
			Id:           0,
			UniversityId: 1,
			Name:         "Lanchonete DCE",
		}
		err := RestaurantRepository.AddNew(&rest)
		if err != nil {
			t.Fatal("Error is not nil")
		}
		if !rest.IsValid() {
			t.Fatal("Restaurant is not valid")
		}
	})

	t.Run("Add new restaurant with invalid name", func(t *testing.T) {
		rest := models.Restaurant{
			Id:           0,
			UniversityId: 1,
			Name:         "",
		}
		err := RestaurantRepository.AddNew(&rest)
		if err.Code() != RestaurantRepository.INV_NAME {
			t.Fatal("Error is nil")
		}
		if rest.IsValid() {
			t.Fatal("Restaurant is  valid")
		}
	})

	t.Run("Add new restaurant with invalid Uni", func(t *testing.T) {
		rest := models.Restaurant{
			Id:           0,
			UniversityId: 0,
			Name:         "Hello There",
		}
		err := RestaurantRepository.AddNew(&rest)
		if err.Code() != RestaurantRepository.INV_UNI {
			t.Fatal("Error is nil")
		}
		if rest.IsValid() {
			t.Fatal("Restaurant is valid")
		}
	})
}

func TestUniversityRepository(t *testing.T) {
	// Change to root dir
	os.Chdir("..")

	// Create database for this test
	stamp := time.Now().Unix()
	// Use memory mode (does not create local database file)
	database.InitSQLite(fmt.Sprintf("file:db-%d.db?mode=memory", stamp))
	defer database.CloseDB()

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

	t.Run("Add new University", func(t *testing.T) {
		uni := models.University{
			Id:   0,
			Name: "PUCPR campus curitiba",
		}
		err := UniversityRepository.AddNew(&uni)
		if err != nil {
			t.Fatal("Error is not nil")
		}
		if !uni.IsValid() {
			t.Fatal("Restaurant is not valid")
		}
	})

	t.Run("Add invalid University", func(t *testing.T) {
		uni := models.University{
			Id:   0,
			Name: "",
		}
		err := UniversityRepository.AddNew(&uni)
		if err.Code() != UniversityRepository.INV_NAME {
			t.Fatal("Error is nil")
		}
		if uni.IsValid() {
			t.Fatal("Restaurant is valid")
		}
	})
}

func TestUserReferencerepository(t *testing.T) {
	// Change to root dir
	os.Chdir("..")

	// Create database for this test
	stamp := time.Now().Unix()
	// Use memory mode (does not create local database file)
	database.InitSQLite(fmt.Sprintf("file:db-%d.db?mode=memory", stamp))
	defer database.CloseDB()

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

	t.Run("Add valid UserReference item", func(t *testing.T) {
		obj := models.UserReference{
			Id:           0,
			UniversityId: 1,
			RestaurantId: 0,
			UserId:       1,
		}
		err := UserReferenceRepository.AddNew(&obj)
		if err != nil {
			t.Fatal(err.Error())
		}
		if !obj.IsValid() {
			t.Fatal("Obj is invalid")
		}
	})

	t.Run("Add invalid UserReference item", func(t *testing.T) {
		obj := models.UserReference{
			Id:           0,
			UniversityId: 0,
			RestaurantId: 0,
			UserId:       1,
		}
		err := UserReferenceRepository.AddNew(&obj)
		if err.Code() != UserReferenceRepository.INV_CONFIG {
			t.Fatal(err.Error())
		}
		if obj.IsValid() {
			t.Fatal("Obj is valid")
		}
	})

	t.Run("Add invalid UserReference item", func(t *testing.T) {
		obj := models.UserReference{
			Id:           0,
			UniversityId: 1,
			RestaurantId: 1,
			UserId:       1,
		}
		err := UserReferenceRepository.AddNew(&obj)
		if err.Code() != UserReferenceRepository.INV_CONFIG {
			t.Fatal(err.Error())
		}
		if obj.IsValid() {
			t.Fatal("Obj is valid")
		}
	})

	t.Run("Add invalid UserReference item", func(t *testing.T) {
		obj := models.UserReference{
			Id:           0,
			UniversityId: 1,
			RestaurantId: 0,
			UserId:       0,
		}
		err := UserReferenceRepository.AddNew(&obj)
		if err.Code() != UserReferenceRepository.INV_USR {
			t.Fatal(err.Error())
		}
		if obj.IsValid() {
			t.Fatal("Obj is valid")
		}
	})
}
