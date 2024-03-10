package repository

import (
	"errors"

	"threebaristas.com/purple/app/models"
)

type CategoriesRepositoryImpl struct {
	CategoryList []*models.Category
}

func NewCategoriesRepositoryImpl() *CategoriesRepositoryImpl {
	return &CategoriesRepositoryImpl{
		CategoryList: models.GetCategoryListExample(),
	}
}

func (r *CategoriesRepositoryImpl) GetCategoryByID(id int64) (*models.Category, error) {
  list := r.CategoryList
  for _, category := range list {
    println(category.ID, category.Name)
  }
  for _, category := range list {
    println(category.ID, category.Name)
    if category.ID == id {
      return category, nil
    }
  }
  return nil, errors.New("Not found")
}
