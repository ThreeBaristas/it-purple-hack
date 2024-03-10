package repository

import (
	"threebaristas.com/purple/app/models"
)

type CategoriesRepositoryImpl struct {
}

func NewCategoriesRepositoryImpl() *CategoriesRepositoryImpl {
	return &CategoriesRepositoryImpl{}
}

func (r *CategoriesRepositoryImpl) GetCategoryByID(id int64) (*models.Category, error) {

	//For Ex.
	category := &models.Category{
		ID:   1,
		Name: "Личные вещи",
	}

	return category, nil
}
