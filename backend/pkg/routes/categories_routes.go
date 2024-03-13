package routes

import (
	"github.com/gofiber/fiber/v2"
	"threebaristas.com/purple/app/controllers"
	"threebaristas.com/purple/app/repository"
)

func CategoriesRoutes(a *fiber.App, repo *repository.CategoriesRepositoryImpl) {
	CategoriesController := controllers.NewCategoriesController(repo)
	a.Get("/api/v1/categories", CategoriesController.GetCategoriesBySearch)
	route := a.Group("/api/v1/categories")
	route.Get("/:id", CategoriesController.GetCategoryByID)
}
