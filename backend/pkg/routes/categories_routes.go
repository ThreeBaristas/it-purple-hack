package routes

import (
	"github.com/gofiber/fiber/v2"
	"threebaristas.com/purple/app/controllers"
	"threebaristas.com/purple/app/repository"
)

func CategoriesRoutes(a *fiber.App) {
	CategoriesController := controllers.CategoriesController{
		CategoriesRepo: repository.NewCategoriesRepositoryImpl(),
	}
	route := a.Group("/api/v1/categories")
	route.Get("/:id", CategoriesController.GetCategoryByID)
}
