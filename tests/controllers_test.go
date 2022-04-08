package tests

import (
	"TapTake-server/app"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRestaurant(t *testing.T) {
	router := app.Init()

	t.Run("Get all restaurants", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/restaurant", nil)
		res := httptest.NewRecorder()
		router.ServeHTTP(res, req)
		result := res.Result()
		defer result.Body.Close()
		data, err := ioutil.ReadAll(result.Body)

		if err != nil {
			t.Errorf("Expected error to be nil, got %v", err)
		}
		if string(data) != "Restaurantes" {
			t.Errorf("Expected data to be Restaurantes, got %s", data)
		}
	})
}
