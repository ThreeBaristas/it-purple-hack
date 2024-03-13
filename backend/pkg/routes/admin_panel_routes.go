package routes

import (
	"github.com/gofiber/fiber/v2"
	"threebaristas.com/purple/app/controllers"
	"threebaristas.com/purple/app/core/services"
	"threebaristas.com/purple/app/repository"
)

func AdminPanelRoutes(a *fiber.App, cR *repository.CategoriesRepository, lR *repository.LocationsRepository, pR *repository.PriceRepository, storage *repository.MatricesMappingStorage) {
	service := services.NewPriceService(cR, lR, pR, storage)
	controller := controllers.NewAdminController(&service, storage)
	route := a.Group("/api/v1/admin")
	route.Get("/price", controller.GetPrice)
	route.Put("/price", controller.SetPrice)
	route.Delete("/price", controller.DeletePrice)
	route.Get("/rules", controller.GetRules)
	route.Post("/storage", controller.SetUpStorage)
	route.Get("/storage", controller.GetStorage)
}
