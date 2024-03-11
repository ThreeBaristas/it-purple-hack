package routes

import (
	"github.com/gofiber/fiber/v2"
	"threebaristas.com/purple/app/controllers"
	"threebaristas.com/purple/app/core/services"
	"threebaristas.com/purple/app/repository"
)

func AdminPanelRoutes(a *fiber.App, cR *repository.CategoriesRepository, lR *repository.LocationsRepository, pR *repository.PriceRepository) {
	service := services.NewPriceService(cR, lR, pR)
	controller := controllers.NewAdminController(&service)
	route := a.Group("/api/v1/admin")
	route.Get("/price", controller.GetPrice)
	route.Put("/price", controller.SetPrice)
	route.Delete("/price", controller.DeletePrice)
	route.Get("/rules", controller.GetRules)
}
