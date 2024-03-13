package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"threebaristas.com/purple/app/core/services"
	"threebaristas.com/purple/app/repository"
)

type PriceController struct {
	priceService    *services.PriceService
	segmentsService *services.GetUserSegmentsService
  mapper *repository.MatricesMappingStorage
}

func NewPriceController(
	service *services.PriceService,
	segmentsService *services.GetUserSegmentsService,
  mapper *repository.MatricesMappingStorage,
) PriceController {
	return PriceController{
		priceService:    service,
		segmentsService: segmentsService,
    mapper: mapper,
	}
}

// GetPrice func gets a price for given category_id & location_id & user_id
//
//	@Tags		admin
//	@Produce	json
func (a *PriceController) GetPrice(c *fiber.Ctx) error {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	categoryId, err := strconv.ParseInt(c.Query("category_id", "NULL"), 10, 64)

	if err != nil {
		logger.Error("Could not parse category_id", zap.Error(err))
		c.SendStatus(400)
		return c.SendString("Error! category_id is not a number")
	}

	locationId, err := strconv.ParseInt(c.Query("location_id", "NULL"), 10, 64)
	if err != nil {
		logger.Error("Could not parse location_id", zap.Error(err))
		c.SendStatus(400)
		return c.SendString("Error! location_id is not a number")
	}

	userId, err := strconv.ParseInt(c.Query("user_id", "-1"), 10, 64)
	if err != nil {
		logger.Error("Could not parse user_id", zap.Error(err))
		return c.SendString("Error! user_id is not a number")
	}

	logger.Info("Handling /price request", zap.Int64("categoryId", categoryId), zap.Int64("locationId", locationId), zap.Int64("userId", userId))
	segments, err := a.segmentsService.GetSegments(userId)
	if err != nil {
		c.SendStatus(500)
		logger.Error("Could not get user segments", zap.Error(err))
		return c.SendString("Error. could not get segments" + err.Error())
	}

	resp, err := a.priceService.GetPrice(locationId, categoryId, segments)
	if err != nil {
		c.SendStatus(500)
		logger.Error("Could not compute price", zap.Error(err))
		return c.SendString("Error. could not find price. " + err.Error())
	}

  segment_id, _ := (*a.mapper).GetSegmentByMatrix(resp.MatrixId)

	return c.JSON(fiber.Map{
		"price":       resp.Price,
		"category_id": resp.CategoryId,
		"location_id": resp.LocationId,
		"matrix_id":   resp.MatrixId,
		"segment_id":  segment_id,
	})
}
