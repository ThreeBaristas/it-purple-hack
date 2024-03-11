package test

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"threebaristas.com/purple/app/controllers"
	"threebaristas.com/purple/app/models"
	"threebaristas.com/purple/app/repository"
)

func TestGetCategoryByID(t *testing.T) {
	mockRepo := repository.NewCategoriesRepositoryImpl()
	controller := controllers.NewCategoriesController(&mockRepo)

	app := fiber.New()
	app.Get("/categories/:id", controller.GetCategoryByID)

	t.Run("Get existing category by ID", func(t *testing.T) {
		expectedCategory := &models.Category{
			ID:   8,
			Name: "Ноутбуки",
		}
		categoryID := strconv.FormatInt(expectedCategory.ID, 10)

		req := httptest.NewRequest(http.MethodGet, "/categories/"+categoryID, nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		var returnedCategory models.Category
		err = json.NewDecoder(resp.Body).Decode(&returnedCategory)
		assert.NoError(t, err)
		assert.Equal(t, expectedCategory, &returnedCategory)
	})
}

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
}

func TestGetLocationsBySearch(t *testing.T) {
	repo := repository.NewLocationsRepositoryImpl()
	controller := controllers.NewLocationsController(&repo)

	app := fiber.New()
	app.Get("/locations/search", controller.GetLocationsBySearch)

	t.Run("GetLocationsBySearch_ValidSearch_Success", func(t *testing.T) {
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

	t.Run("GetLocationsBySearch_EmptySearch_BadRequest", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/locations/search", nil)
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})
}
