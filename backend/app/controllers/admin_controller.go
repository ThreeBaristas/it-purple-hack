package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"threebaristas.com/purple/app/core/services"
	"threebaristas.com/purple/app/repository"
)

type AdminController struct {
	service *services.PriceService
	// REFACTOR: add some kind of service
	storageRepo *repository.MatricesMappingStorage
}

func NewAdminController(
	service *services.PriceService,
	storageRepo *repository.MatricesMappingStorage,
) AdminController {
	return AdminController{
		service:     service,
		storageRepo: storageRepo,
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

	matrixId, err := strconv.ParseInt(c.Query("matrix_id", "0"), 10, 64)
	if err != nil {
		c.SendStatus(400)
		return c.SendString("Error! matrix_id is not a number")
	}

	logger.Info("Handling /admin/price request", zap.Int64("categoryId", categoryId), zap.Int64("locationId", locationId), zap.Int64("matrixId", matrixId))
	resp, err := a.service.GetPrice(locationId, categoryId, []int64{matrixId}, nil)
	if err != nil {
		c.SendStatus(500)
		logger.Error("Could not compute price", zap.Error(err))
		return c.SendString("Error. could not find price. " + err.Error())
	}

	return c.JSON(fiber.Map{
		"price":       resp.Price,
		"category_id": resp.CategoryId,
		"location_id": resp.LocationId,
		"matrix_id":   resp.MatrixId,
	})
}

func (a *AdminController) SetPrice(c *fiber.Ctx) error {
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

	matrixId, err := strconv.ParseInt(c.Query("matrix_id", "0"), 10, 64)
	if err != nil {
		logger.Error("Could not parse matrix_id", zap.Error(err))
		return c.SendString("Error! matrix_id is not a number")
	}

	price, err := strconv.ParseInt(c.Query("price", "NULL"), 10, 64)
	if err != nil {
		logger.Error("Could not parse price", zap.Error(err))
		return c.SendString("Error! price is not a number")
	}

	resp, err := a.service.SetPrice(locationId, categoryId, matrixId, price)
	if err != nil {
		logger.Error("Could not set price", zap.Error(err))
		return c.SendString("Error! could not set price")
	}

	return c.JSON(resp)
}

func (a *AdminController) DeletePrice(c *fiber.Ctx) error {
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

	matrixId, err := strconv.ParseInt(c.Query("matrix_id", "0"), 10, 64)
	if err != nil {
		logger.Error("Could not parse matrix_id", zap.Error(err))
		return c.SendString("Error! matrix_id is not a number")
	}

	resp, err := a.service.DeletePrice(locationId, categoryId, matrixId)
	if err != nil {
		logger.Error("Could not delete price", zap.Error(err))
		return c.SendString("Error! could not delete price")
	}

	return c.JSON(resp)
}

func (a *AdminController) GetRules(c *fiber.Ctx) error {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	page, err := strconv.ParseInt(c.Query("page", "0"), 10, 64)
	if err != nil {
		logger.Error("Could not parse page", zap.Error(err))
		c.SendStatus(400)
		return c.SendString("Error! page is not a number")
	}

	pageSize, err := strconv.ParseInt(c.Query("pageSize", "10"), 10, 64)
	if err != nil {
		logger.Error("Could not parse pageSize", zap.Error(err))
		c.SendStatus(400)
		return c.SendString("Error! pageSize is not a number")
	}

	rules, err := a.service.GetRules(services.GetPricesRequest{Page: page, PageSize: int32(pageSize)})
	if err != nil {
		logger.Error("Could not get rules", zap.Error(err))
		c.SendStatus(500)
		return c.SendString("Error! could not get rules")
	}
	logger.Info("Got all rules from db", zap.Int64("totalPages", int64(rules.TotalPages)))
	var dtosArray []RuleDTO
	for _, val := range rules.Data {
		dtosArray = append(dtosArray, RuleDTO{
			Location: LocationDTO{
				ID:   val.Location.ID,
				Name: val.Location.Name,
			},
			Category: CategoryDTO{
				ID:   val.Category.ID,
				Name: val.Category.Name,
			},
			MatrixId: val.Matrix,
			Price:    val.Price,
		})
	}
	return c.JSON(GetRulesResponse{
		Data:      dtosArray,
		TotalPage: int64(rules.TotalPages),
		Page:      page,
		PageSize:  pageSize,
	})
}

type RuleDTO struct {
	Location LocationDTO `json:"location"`
	Category CategoryDTO `json:"category"`
	MatrixId int64       `json:"matrix_id"`
	Price    int64       `json:"price"`
}

type GetRulesResponse struct {
	Data      []RuleDTO `json:"data"`
	TotalPage int64     `json:"totalPages"`
	Page      int64     `json:"page"`
	PageSize  int64     `json:"pageSize"`
}

func (a *AdminController) SetUpStorage(c *fiber.Ctx) error {
	var payload repository.SetUpStorageRequest

	if err := c.BodyParser(&payload); err != nil {
		c.SendStatus(400)
		return err
	}

	return (*a.storageRepo).SetUpStorage(&payload)
}

func (a *AdminController) GetStorage(c *fiber.Ctx) error {
	storage, err := (*a.storageRepo).GetStorage()
	if err != nil {
		c.SendStatus(500)
		return err
	}
	filteredDiscounts := []repository.DiscountMappingDTO{}
	for _, dto := range storage.Discounts {
		if dto.MatrixId != storage.BaselineMatrix {
			filteredDiscounts = append(filteredDiscounts, dto)
		}
	}
	storage.Discounts = filteredDiscounts
	return c.JSON(storage)
}
