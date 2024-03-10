package repository

import (
	"threebaristas.com/purple/app/models"
)

type CategoriesRepositoryImpl struct {
	CategoryList models.CategoryList
}

func NewCategoriesRepositoryImpl() *CategoriesRepositoryImpl {
	return &CategoriesRepositoryImpl{
		CategoryList: *models.GetCategoryListExample(),
	}
}

func (r *CategoriesRepositoryImpl) GetCategoryByID(id int64) (*models.Category, error) {
	category, err := r.CategoryList.GetCategoryByID(id)
	if err != nil {
		return nil, err
	}
	return category, nil
}
