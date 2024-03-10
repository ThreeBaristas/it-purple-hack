package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"strconv"
	"threebaristas.com/purple/app/repository"
)

type CategoriesController struct {
	CategoriesRepo repository.CategoriesRepository
}

func NewCategoriesController(categoriesRepo repository.CategoriesRepository) *CategoriesController {
	return &CategoriesController{CategoriesRepo: categoriesRepo}
}

// GetCategoryByID func gets a category by id
//
//	@Tags		admin
//	@Produce	json
func (c *CategoriesController) GetCategoryByID(ctx *fiber.Ctx) error {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	categoryID, err := strconv.ParseInt(ctx.Params("id"), 10, 64)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest)
		return ctx.SendString("Error! category_id is not a number")
	}

	logger.Info("Handling /categories/{id} request", zap.Int64("categoryID", categoryID))
	category, err := c.CategoriesRepo.GetCategoryByID(categoryID)
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError)
		return ctx.SendString("Error! Failed to get category")
	}
	return ctx.JSON(fiber.Map{
		"id":   category.ID,
		"name": category.Name,
	})
}
