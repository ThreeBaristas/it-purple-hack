package repository

import (
	"errors"
	"strings"

	"threebaristas.com/purple/app/models"
)

type LocationsRepository interface {
	GetLocationByID(id int64) (*models.Location, error)
	GetByString(s string, n int) ([]*models.Location, error)
}
type LocationsRepositoryImpl struct {
	LocationRoot *models.Location
}

func NewLocationsRepositoryImpl() LocationsRepository {
	return &LocationsRepositoryImpl{
		LocationRoot: models.GetLocationTreeExample(),
	}
}

func (r *LocationsRepositoryImpl) GetLocationByID(id int64) (*models.Location, error) {
	list := r.LocationRoot.Traverse()
	for _, category := range list {
		if category.ID == id {
			return category, nil
		}
	}
	return nil, errors.New("Not found")
}

func (r *LocationsRepositoryImpl) GetByString(s string, max int) ([]*models.Location, error) {
	list := r.LocationRoot.Traverse()
	var filtered []*models.Location
	for _, category := range list {
		if strings.Contains(strings.ToLower(category.Name), strings.ToLower(s)) {
			filtered = append(filtered, category)
		}
		if len(filtered) >= max {
			break
		}
	}
	return filtered, nil
}
