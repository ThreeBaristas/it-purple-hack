package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"strconv"
	"threebaristas.com/purple/app/core/services"
)

type AdminController struct {
	service *services.PriceService
}

func NewAdminController(
	service *services.PriceService,
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

	segmentId, err := strconv.ParseInt(c.Query("segment_id", "0"), 10, 64)
	if err != nil {
    logger.Error("Could not parse segment_id", zap.Error(err))
		return c.SendString("Error! segment_id is not a number")
	}


	price, err := strconv.ParseInt(c.Query("price", "NULL"), 10, 64)
	if err != nil {
    logger.Error("Could not parse price", zap.Error(err))
		return c.SendString("Error! price is not a number")
	}

  resp, err := a.service.SetPrice(locationId, categoryId, segmentId, price)
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

	segmentId, err := strconv.ParseInt(c.Query("segment_id", "0"), 10, 64)
	if err != nil {
    logger.Error("Could not parse segment_id", zap.Error(err))
		return c.SendString("Error! segment_id is not a number")
	}

  resp, err := a.service.DeletePrice(locationId, categoryId, segmentId)
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
        ID: val.Location.ID,
        Name: val.Location.Name,
      },
      Category: CategoryDTO{
        ID: val.Category.ID,
        Name: val.Category.Name,
      },
      MatrixId: val.Segment,
      Price: val.Price,
    })
  }
  return c.JSON(GetRulesResponse {
    Data: dtosArray,
    TotalPage: int64(rules.TotalPages),
  });
}

type RuleDTO struct {
  Location LocationDTO `json:"location"`
  Category CategoryDTO `json:"category"`
  MatrixId int64 `json:"segment"`
  Price int64 `json:"price"`
}

type GetRulesResponse struct {
  Data []RuleDTO `json:"data"`
  TotalPage int64 `json:"totalPages"`
}
