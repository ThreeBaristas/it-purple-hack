package admin

import "threebaristas.com/purple/app/repository"

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

func (a *AdminService) GetPrice(locationId int64, categoryId int64) int64 {
	return 42
}
