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
		if category.ID == id {
			return category, nil
		}
	}
	return nil, errors.New("Not found")
}

func (r *CategoriesRepositoryImpl) GetByString(s string, max int) ([]*models.Category, error) {
	list := r.CategoryList
	var filtered []*models.Category
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
