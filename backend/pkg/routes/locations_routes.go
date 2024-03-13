package routes

import (
	"github.com/gofiber/fiber/v2"
	"threebaristas.com/purple/app/controllers"
	"threebaristas.com/purple/app/repository"
)

func LocationsRoutes(a *fiber.App, repo *repository.LocationsRepositoryImpl) {
	controller := controllers.NewLocationsController(repo)
	a.Get("/api/v1/locations", controller.GetLocationsBySearch)
	route := a.Group("/api/v1/locations")
	route.Get("/:id", controller.GetLocationByID)
}
