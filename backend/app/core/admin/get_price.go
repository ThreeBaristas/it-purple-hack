package admin

import (
	"errors"
	"threebaristas.com/purple/app/models"
	"threebaristas.com/purple/app/repository"
)

type AdminService struct {
	categoriesRepo *repository.CategoriesRepository
	locationsRepo  *repository.LocationsRepository
	priceRepo      *repository.PriceRepository
}

func NewAdminService(
	categoriesRepo *repository.CategoriesRepository,
	locationsRepo *repository.LocationsRepository,
	priceRepo *repository.PriceRepository,
) AdminService {
	return AdminService{
		categoriesRepo: categoriesRepo,
		locationsRepo:  locationsRepo,
		priceRepo:      priceRepo,
	}
}

func (a *AdminService) GetPrice(locationId int64, categoryId int64, segmentsIds []int64) (*repository.GetPriceResponse, error) {

	location, _ := (*a.locationsRepo).GetLocationByID(locationId)
	if location == nil {
		return nil, errors.New("Location not found")
	}
	category, _ := (*a.categoriesRepo).GetCategoryByID(categoryId)
	if category == nil {
		return nil, errors.New("Category not found")
	}

	// Массив локаций на пути от стартовой до рута
	var locations []*models.Location
	var categories []*models.Category

	locationCur := location
	for locationCur != nil {
		locations = append(locations, locationCur)
		locationCur = locationCur.Parent
	}
	categoryCur := category
	for categoryCur != nil {
		categories = append(categories, categoryCur)
		categoryCur = categoryCur.Parent
	}

	req := formBatchRequest(locations, categories)
	var matricesIds []int64
	for _, value := range segmentsIds {
		matricesIds = append(matricesIds, a.SegmentToMatrixId(value))
	}
	req.Matrices = matricesIds

	response, err := (*a.priceRepo).GetPricesBatch(req)
	if err != nil {
		return nil, err
	}

	firstGoodNode := findFirstNode(locations, categories, response)

	return firstGoodNode, nil
}

func (a *AdminService) SegmentToMatrixId(segmentId int64) int64 {
	return segmentId
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
 * @param `response` - array of prices for pairs of `category` and `location`
 **/
func findFirstNode(locations []*models.Location, categories []*models.Category, response []repository.GetPriceResponse) *repository.GetPriceResponse {
	for _, loc := range locations {
		for _, cat := range categories {
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
