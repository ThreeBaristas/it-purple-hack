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
	asMap        map[int64]*models.Location
	AsList       []*models.Location
}

func NewLocationsRepositoryImpl() *LocationsRepositoryImpl {
	root := models.GetLocationTreeExample()
	asList := root.Traverse()
	asMap := make(map[int64]*models.Location)
	for _, node := range asList {
		asMap[node.ID] = node
	}

	return &LocationsRepositoryImpl{
		LocationRoot: root,
		AsList:       asList,
		asMap:        asMap,
	}
}

/**
 * Returns a location with such id or an error if no location exists. It works in O(1).
 **/
func (r *LocationsRepositoryImpl) GetLocationByID(id int64) (*models.Location, error) {
	res, ok := r.asMap[id]
	if ok {
		return res, nil
	}
	return nil, errors.New("Not found")
}

/**
 * Returns locations which names contain given string. It works in O(n)
 **/
func (r *LocationsRepositoryImpl) GetByString(s string, max int) ([]*models.Location, error) {
	var filtered []*models.Location
	for _, category := range r.AsList {
		if strings.Contains(strings.ToLower(category.Name), strings.ToLower(s)) {
			filtered = append(filtered, category)
		}
		if len(filtered) >= max {
			break
		}
	}
	return filtered, nil
}
