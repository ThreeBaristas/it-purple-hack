package test

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"threebaristas.com/purple/app/controllers"
	"threebaristas.com/purple/app/models"
	"threebaristas.com/purple/app/repository"
)

func TestGetCategoryByID(t *testing.T) {
	mockRepo := repository.NewCategoriesRepositoryImpl()
	controller := controllers.NewCategoriesController(mockRepo)

	app := fiber.New()
	app.Get("/categories/:id", controller.GetCategoryByID)

	t.Run("Get existing category by ID", func(t *testing.T) {
		expectedCategory := &models.Category{
			ID:   1,
			Name: "Личные вещи",
		}
		categoryID := "1"

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
