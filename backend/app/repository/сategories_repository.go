package repository

import "threebaristas.com/purple/app/models"

type CategoriesRepository interface {
	GetCategoryByID(id int64) (*models.Category, error)
}
