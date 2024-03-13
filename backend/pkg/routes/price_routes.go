package routes

import (
	"github.com/gofiber/fiber/v2"
	"threebaristas.com/purple/app/controllers"
	"threebaristas.com/purple/app/core/services"
	"threebaristas.com/purple/app/repository"
)

func PriceRoutes(a *fiber.App, cR *repository.CategoriesRepository, lR *repository.LocationsRepository, pR *repository.PriceRepository, storage *repository.MatricesMappingStorage) {
	service := services.NewPriceService(cR, lR, pR, storage)
	segmentsService := services.NewGetUserSegmentsService()
	controller := controllers.NewPriceController(&service, segmentsService, storage)
	a.Get("/api/v1/price", controller.GetPrice)
}
