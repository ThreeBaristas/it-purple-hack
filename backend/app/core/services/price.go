package services

import (
	"errors"
	"threebaristas.com/purple/app/models"
	"threebaristas.com/purple/app/repository"
)

type PriceService struct {
	categoriesRepo *repository.CategoriesRepository
	locationsRepo  *repository.LocationsRepository
	priceRepo      *repository.PriceRepository
	storage        *repository.MatricesMappingStorage
}

type PriceRule struct {
	Location *models.Location `json:"location"`
	Category *models.Category `json:"category"`
	Segment  int64            `json:"segment"`
	Price    int64            `json:"price"`
}

type GetRulesResponse struct {
	Data       []PriceRule `json:"data"`
	TotalPages int32       `json:"totalPages"`
}

type GetPricesRequest struct {
	Page     int64
	PageSize int32
}

func NewPriceService(
	categoriesRepo *repository.CategoriesRepository,
	locationsRepo *repository.LocationsRepository,
	priceRepo *repository.PriceRepository,
	storage *repository.MatricesMappingStorage,
) PriceService {
	return PriceService{
		categoriesRepo: categoriesRepo,
		locationsRepo:  locationsRepo,
		priceRepo:      priceRepo,
		storage:        storage,
	}
}

func (a *PriceService) SetPrice(locationId int64, categoryId int64, segmentsId int64, price int64) (*repository.GetPriceResponse, error) {
	matrixId, ok := (*a.storage).SegmentToMatrix(segmentsId)
	if !ok {
		return nil, errors.New("Segment not found")
	}
	return (*a.priceRepo).SetPrice(locationId, categoryId, matrixId, price)
}

func (a *PriceService) DeletePrice(locationId int64, categoryId int64, segmentId int64) (bool, error) {
	matrixId, ok := (*a.storage).SegmentToMatrix(segmentId)
	if !ok {
		return false, errors.New("Segment not found")
	}
	return (*a.priceRepo).DeletePrice(locationId, categoryId, matrixId)
}

func (a *PriceService) GetRules(req GetPricesRequest) (*GetRulesResponse, error) {
	res, totalPages, err := (*a.priceRepo).GetRules(req.PageSize, req.Page)
	if err != nil {
		return nil, err
	}
	var res1 []PriceRule
	for _, r := range res {
		location, err := (*a.locationsRepo).GetLocationByID(r.LocationId)
		if err != nil {
			return nil, err
		}
		category, err := (*a.categoriesRepo).GetCategoryByID(r.CategoryId)
		if err != nil {
			return nil, err
		}

		res1 = append(res1, PriceRule{
			Location: location,
			Category: category,
			Segment:  r.MatrixId,
			Price:    r.Price,
		})
	}
	return &GetRulesResponse{
		Data:       res1,
		TotalPages: int32(totalPages),
	}, nil
}

func (a *PriceService) GetPrice(locationId int64, categoryId int64, segmentsIds []int64) (*repository.GetPriceResponse, error) {
  // O(1)
	location, _ := (*a.locationsRepo).GetLocationByID(locationId)
	if location == nil {
		return nil, errors.New("Location not found")
	}
  // O(1)
	category, _ := (*a.categoriesRepo).GetCategoryByID(categoryId)
	if category == nil {
		return nil, errors.New("Category not found")
	}

	// Массив локаций на пути от стартовой до рута
	var locations []*models.Location
	var categories []*models.Category

  // O(h_1)
	locationCur := location
	for locationCur != nil {
		locations = append(locations, locationCur)
		locationCur = locationCur.Parent
	}
  // O(h_2)
	categoryCur := category
	for categoryCur != nil {
		categories = append(categories, categoryCur)
		categoryCur = categoryCur.Parent
	}

  // O(h_1 + h_2)
	req := formBatchRequest(locations, categories)
	var matricesIds []int64
	for _, value := range segmentsIds {
		matrixId, ok := (*a.storage).SegmentToMatrix(value)
		if !ok {
			return nil, errors.New("Segment not found")
		}
		matricesIds = append(matricesIds, matrixId)
	}
	req.Matrices = matricesIds

  // Response has a length of O(h^2)
	response, err := (*a.priceRepo).GetPricesBatch(req)
	if err != nil {
		return nil, err
	}

  // O(h^2) (!!!)
	firstGoodNode := a.findFirstNode(locations, categories, response)

	return firstGoodNode, nil
}

func formBatchRequest(locations []*models.Location, categories []*models.Category) *repository.GetPriceRequest {
	var locationsIds []int64
	var categoryIds []int64
	// location pointer
	for _, loc := range locations {
		locationsIds = append(locationsIds, loc.ID)
	}
	for _, category := range categories {
		categoryIds = append(categoryIds, category.ID)
	}
	return &repository.GetPriceRequest{
		LocationId: locationsIds,
		CategoryId: categoryIds,
		Matrices:   nil,
	}
}

/** Returns a first node from `response` array that is first from `RoadUpSearch` algortihm's perspectife
 * @param `locations` - array of locations from `location_start` to `ROOT`
 * @param `categories` - array of categories from `category_start` to `ROOT`
 * @param `response` - array of prices for pairs of `category` and `location`. This array is ordered by `MatrixId` in descending order
 **/
func (a *PriceService) findFirstNode(locations []*models.Location, categories []*models.Category, response []repository.GetPriceResponse) *repository.GetPriceResponse {
  // Сначала поднимаеся по категориям, потом по локациям.
  // Поэтому первым делом выбираем пару, у которой локация дальше всего от корня 
  maxLocation, _ := (*a.locationsRepo).GetLocationByID(response[0].LocationId)
  for _, node := range response {
    nodeLocation, _ := (*a.locationsRepo).GetLocationByID(node.LocationId)
    if nodeLocation.DistToRoot > maxLocation.DistToRoot {
      maxLocation = nodeLocation
    }
  }

  maxCategoryMaxLocation, _ := (*a.categoriesRepo).GetCategoryByID(response[0].CategoryId)
  for _, node := range response {
    nodeCategory, _ := (*a.categoriesRepo).GetCategoryByID(node.CategoryId)
    if nodeCategory.DistToRoot > maxCategoryMaxLocation.DistToRoot {
      maxCategoryMaxLocation = nodeCategory
    }
  }

  var answer *repository.GetPriceResponse = nil
  for _, node := range response {
    if node.CategoryId == maxCategoryMaxLocation.ID && node.LocationId == maxLocation.ID {
      answer = &node
      break
    }
  }
  
	return answer
}

func findInResponse(location *models.Location, category *models.Category, response []repository.GetPriceResponse) *repository.GetPriceResponse {
	for _, elem := range response {
		if elem.LocationId == location.ID && elem.CategoryId == category.ID {
			return &elem
		}
	}
	return nil
}
