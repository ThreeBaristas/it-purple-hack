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
	t.Run("GetCategoryByID_InvalidID_BadRequest", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/categories/invalid_id", nil)
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})
	t.Run("GetCategoryByID_NonExistentID_NotFound", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/categories/999", nil)
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	})
}

func TestGetCategoriesBySearch(t *testing.T) {
	expectedCategories := []*models.Category{
		{ID: 8, Name: "Ноутбуки"},
	}
	mockRepo := repository.NewCategoriesRepositoryImpl()
	controller := controllers.NewCategoriesController(&mockRepo)

	app := fiber.New()
	app.Get("/categories/search", controller.GetCategoriesBySearch)

	t.Run("Get existing categories by search query", func(t *testing.T) {

		searchQuery := "Ноутбуки"

		req := httptest.NewRequest(http.MethodGet, "/categories/search?search="+searchQuery, nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		var returnedCategories []*models.Category
		err = json.NewDecoder(resp.Body).Decode(&returnedCategories)
		assert.NoError(t, err)
		assert.ElementsMatch(t, expectedCategories, returnedCategories)
	})

	t.Run("GetCategoriesByEmptySearch_Success", func(t *testing.T) {
		// Создание примера дерева локаций
		locationTree := models.GetCategoryTreeExample()

		req := httptest.NewRequest(http.MethodGet, "/categories/search", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		// Чтение тела ответа
		body, err := io.ReadAll(resp.Body)
		assert.NoError(t, err)

		var locations []map[string]interface{}
		err = json.Unmarshal(body, &locations)
		assert.NoError(t, err)

		// Проверка, что все значения присутствуют в ответе
		for _, expectedLocation := range locationTree.Children {
			found := false
			for _, actualLocation := range locations {
				if expectedLocation.Name == actualLocation["name"].(string) {
					found = true
					break
				}
			}
			assert.True(t, found, "Expected category not found in response: %s", expectedLocation.Name)
		}
	})
}
