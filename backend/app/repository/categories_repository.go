package repository

import (
	"errors"
	"strings"

	"threebaristas.com/purple/app/models"
)

type CategoriesRepository interface {
	GetCategoryByID(id int64) (*models.Category, error)
	GetByString(s string, n int) ([]*models.Category, error)
}
type CategoriesRepositoryImpl struct {
	root   *models.Category
	asList []*models.Category
	asMap  map[int64]*models.Category
}

func NewCategoriesRepositoryImpl() CategoriesRepository {
	root := models.GetCategoryTreeExample()
	asList := root.FindAllByPredicate(func(c *models.Category) bool { return true })
	asMap := make(map[int64]*models.Category)
	for _, node := range asList {
		asMap[node.ID] = node
	}

	return &CategoriesRepositoryImpl{
		root:   root,
		asList: asList,
		asMap:  asMap,
	}
}

func (r *CategoriesRepositoryImpl) GetCategoryByID(id int64) (*models.Category, error) {
	res, ok := r.asMap[id]
	if !ok {
		return nil, errors.New("Not found")
	}
	return res, nil
}

func (r *CategoriesRepositoryImpl) GetByString(s string, max int) ([]*models.Category, error) {
	var ans []*models.Category
	for _, node := range r.asList {
		if strings.Contains(strings.ToLower(node.Name), strings.ToLower(s)) {
			ans = append(ans, node)

			if len(ans) >= max {
				break
			}
		}
	}
	return ans, nil
}
