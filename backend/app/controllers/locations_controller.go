package controllers

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"threebaristas.com/purple/app/repository"
)

type LocationsController struct {
	repo *repository.LocationsRepository
}

func NewLocationsController(locationsRepo *repository.LocationsRepository) *LocationsController {
	return &LocationsController{repo: locationsRepo}
}

type LocationDTO struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// GetLocationByID func gets a category by id
//
//	@Tags		admin
//	@Produce	json
func (c *LocationsController) GetLocationByID(ctx *fiber.Ctx) error {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	categoryID, err := strconv.ParseInt(ctx.Params("id"), 10, 64)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest)
		return ctx.SendString("Error! category_id is not a number")
	}

	logger.Info("Handling /locations/{id} request", zap.Int64("categoryID", categoryID))
	category, err := (*c.repo).GetLocationByID(categoryID)
	if err != nil {
		// Проверяем, если ошибка не равна strconv.ErrRange,
		// которая указывает на то, что парсинг не удался из-за неправильного значения.
		if err != nil && !errors.Is(err, strconv.ErrRange) {
			ctx.Status(fiber.StatusNotFound)
			return ctx.SendString("Error! Category not found")
		}
		// В противном случае, возвращаем статус код 500
		ctx.Status(fiber.StatusInternalServerError)
		return ctx.SendString("Error! Failed to get category")
	}
	return ctx.JSON(fiber.Map{
		"id":   category.ID,
		"name": category.Name,
	})
}

func (c *LocationsController) GetLocationsBySearch(ctx *fiber.Ctx) error {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	search := ctx.Query("search")

	data, err := (*c.repo).GetByString(search, 10)
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError)
		return ctx.SendString("Error!" + err.Error())
	}

	var dtos []*LocationDTO
	for _, category := range data {
		dtos = append(dtos, &LocationDTO{
			ID:   category.ID,
			Name: category.Name,
		})
	}
	return ctx.JSON(dtos)
}
