package test

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"threebaristas.com/purple/app/controllers"
	"threebaristas.com/purple/app/models"
	"threebaristas.com/purple/app/repository"
)

func TestGetLocationByID(t *testing.T) {
	repo := repository.NewLocationsRepositoryImpl()
	controller := controllers.NewLocationsController(&repo)

	app := fiber.New()
	app.Get("/locations/:id", controller.GetLocationByID)

	t.Run("GetLocationByID_ValidID_Success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/locations/1", nil)
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		// Чтение тела ответа
		body, err := io.ReadAll(resp.Body)
		assert.NoError(t, err)

		// Распарсить JSON из тела ответа
		var location map[string]interface{}
		err = json.Unmarshal(body, &location)
		assert.NoError(t, err)

		// Проверить ожидаемые поля в JSON
		expectedID := 1
		actualID := int(location["id"].(float64))
		assert.Equal(t, expectedID, actualID)
		expectedName := "ROOT"
		actualName := location["name"]
		assert.Equal(t, expectedName, actualName)
	})

	t.Run("GetLocationByID_InvalidID_BadRequest", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/locations/invalid_id", nil)
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})

	t.Run("GetLocationByID_NonExistentID_NotFound", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/locations/999", nil)
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	})
}

func TestGetLocationsBySearch(t *testing.T) {
	repo := repository.NewLocationsRepositoryImpl()
	controller := controllers.NewLocationsController(&repo)

	app := fiber.New()
	app.Get("/locations/search", controller.GetLocationsBySearch)

	t.Run("GetLocationsByValidSearch_Success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/locations/search?search=ROOT", nil)
		resp, err := app.Test(req)

		// Чтение тела ответа
		body, err := io.ReadAll(resp.Body)
		assert.NoError(t, err)

		// Распарсить JSON
		var locations []map[string]interface{}
		err = json.Unmarshal(body, &locations)
		assert.NoError(t, err)

		// Проверка полей
		expectedID := 1
		actualID := int(locations[0]["id"].(float64))
		assert.Equal(t, expectedID, actualID)

		expectedName := "ROOT"
		actualName := locations[0]["name"].(string)
		assert.Equal(t, expectedName, actualName)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("GetLocationsByEmptySearch_Success", func(t *testing.T) {
		// Создание примера дерева локаций
		locationTree := models.GetLocationTreeExample()

		req := httptest.NewRequest(http.MethodGet, "/locations/search", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		// Чтение тела ответа
		body, err := io.ReadAll(resp.Body)
		assert.NoError(t, err)

		var locations []map[string]interface{}
		err = json.Unmarshal(body, &locations)
		assert.NoError(t, err)

		// Проверка, что все 5 значений присутствуют в ответе
		for _, expectedLocation := range locationTree.Children {
			found := false
			for _, actualLocation := range locations {
				if expectedLocation.Name == actualLocation["name"].(string) {
					found = true
					break
				}
			}
			assert.True(t, found, "Expected location not found in response: %s", expectedLocation.Name)
		}
	})
}
