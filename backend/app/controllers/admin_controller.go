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
	resp, err := a.service.GetPrice(locationId, categoryId, nil)
	if err != nil {
		c.SendStatus(500)
		logger.Error("Could not compute price", zap.Error(err))
		return c.SendString("Error. could not find price. " + err.Error())
	}

	return c.JSON(fiber.Map{
		"price":       resp.Price,
		"category_id": resp.CategoryId,
		"location_id": resp.LocationId,
	})
}
