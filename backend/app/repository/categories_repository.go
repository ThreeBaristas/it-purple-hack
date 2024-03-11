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
  root *models.Category
}

func NewCategoriesRepositoryImpl() CategoriesRepository {
	return &CategoriesRepositoryImpl{
    root: models.GetCategoryTreeExample(),
	}
}

func (r *CategoriesRepositoryImpl) GetCategoryByID(id int64) (*models.Category, error) {
  res := r.root.FindChildById(id)
  if(res == nil) {
    return nil, errors.New("Not found")
  }
  return res, nil
}

func (r *CategoriesRepositoryImpl) GetByString(s string, max int) ([]*models.Category, error) {
  return r.root.FindAllByPredicate(func (c *models.Category) bool {
    return strings.Contains(strings.ToLower(c.Name), strings.ToLower(s))
  }), nil
}
