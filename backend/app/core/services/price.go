package services

import (
	"errors"
	"math/rand"
	"sync"

	"threebaristas.com/purple/app/models"
	"threebaristas.com/purple/app/repository"
)

type PriceService struct {
	categoriesRepo *repository.CategoriesRepositoryImpl
	locationsRepo  *repository.LocationsRepositoryImpl
	priceRepo      *repository.PriceRepository
	storage        *repository.MatricesMappingStorage
}

type PriceRule struct {
	Location *models.Location `json:"location"`
	Category *models.Category `json:"category"`
	Matrix   int64            `json:"matrix_id"`
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
	categoriesRepo *repository.CategoriesRepositoryImpl,
	locationsRepo *repository.LocationsRepositoryImpl,
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

func (a *PriceService) SetPrice(locationId int64, categoryId int64, matrixId int64, price int64) (*repository.GetPriceResponse, error) {
	return (*a.priceRepo).SetPrice(locationId, categoryId, matrixId, price)
}

func (a *PriceService) DeletePrice(locationId int64, categoryId int64, matrixId int64) (bool, error) {
	return (*a.priceRepo).DeletePrice(locationId, categoryId, matrixId)
}

func (a *PriceService) GetRules(req GetPricesRequest) (*GetRulesResponse, error) {
	storage, err := (*a.storage).GetStorage()
	var matrices []int64
	for _, entry := range storage.Discounts {
		matrices = append(matrices, entry.MatrixId)
	}

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
			Matrix:   r.MatrixId,
			Price:    r.Price,
		})
	}
	return &GetRulesResponse{
		Data:       res1,
		TotalPages: int32(totalPages),
	}, nil
}

func (a *PriceService) GetPrice(locationId int64, categoryId int64, matrices []int64) (*repository.GetPriceResponse, error) {
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
	req.Matrices = matrices

	// Response has a length of O(h^2)
	response, err := (*a.priceRepo).GetPricesBatch(req)
	if err != nil {
		return nil, err
	}

	// O(h^4) (!!!)
	firstGoodNode := findFirstNode(locations, categories, response)

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
func findFirstNode(locations []*models.Location, categories []*models.Category, response []repository.GetPriceResponse) *repository.GetPriceResponse {
	// O(h) for this array
	for _, loc := range locations {
		// O(h) for this array
		for _, cat := range categories {
			// O(h^2) to find
			res := findInResponse(loc, cat, response)
			if res != nil {
				return res
			}
		}
	}
	return &response[0]
}

func findInResponse(location *models.Location, category *models.Category, response []repository.GetPriceResponse) *repository.GetPriceResponse {
	for _, elem := range response {
		if elem.LocationId == location.ID && elem.CategoryId == category.ID {
			return &elem
		}
	}
	return nil
}

func (p *PriceService) GenerateRules() {
  var wg sync.WaitGroup

  const step = 10
  for i := 0; i < len(p.locationsRepo.AsList); i += step {
    for j := 0; j < len(p.categoriesRepo.AsList); j += step {
      wg.Add(1)
      var matrixId int64
      matrixId = int64(rand.Intn(200))

      locationId := p.locationsRepo.AsList[i].ID
      categoryId := p.categoriesRepo.AsList[j].ID
      price := rand.Intn(5000)
      go func() { p.SetPrice(locationId, categoryId, matrixId, int64(price))
      wg.Done()
    }()
    }
  }

  wg.Wait()
}
