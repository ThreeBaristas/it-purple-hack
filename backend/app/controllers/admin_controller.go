package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"strconv"
	"threebaristas.com/purple/app/core/admin"
)

type AdminController struct {
	service *admin.AdminService
}

func NewAdminController(
	service *admin.AdminService,
) AdminController {
	return AdminController{
		service: service,
	}
}

// GetPrice func gets a price for given category_id & location_id
//
//	@Tags		admin
//	@Produce	json
func (a *AdminController) GetPrice(c *fiber.Ctx) error {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	categoryId, err := strconv.ParseInt(c.Query("category_id", "NULL"), 10, 64)
	if err != nil {
		c.SendStatus(400)
		return c.SendString("Error! category_id is not a number")
	}
	locationId, err := strconv.ParseInt(c.Query("location_id", "NULL"), 10, 64)
	if err != nil {
		c.SendStatus(400)
		return c.SendString("Error! location_id is not a number")
	}
	logger.Info("Handling /price request", zap.Int64("categoryId", categoryId), zap.Int64("locationId", locationId))
	price := a.service.GetPrice(locationId, categoryId)
	return c.JSON(fiber.Map{
		"price":       price,
		"category_id": categoryId,
		"location_id": locationId,
	})
}
