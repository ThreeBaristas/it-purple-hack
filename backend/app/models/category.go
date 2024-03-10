package models

import (
	"errors"
	"fmt"
)

// Category представляет структуру для хранения информации о категории
type Category struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// CategoryList представляет список категорий
type CategoryList struct {
	categories []*Category
}

// newCategoryList создает новый пустой список категорий
func newCategoryList() *CategoryList {
	return &CategoryList{}
}

// addCategory добавляет новую категорию в список
func (cl *CategoryList) addCategory(id int64, name string) {
	category := &Category{
		ID:   id,
		Name: name,
	}
	cl.categories = append(cl.categories, category)
}

// printCategories выводит информацию о всех категориях в списке
func (cl *CategoryList) printCategories() {
	fmt.Println("Categories:")
	for _, category := range cl.categories {
		fmt.Printf("ID: %d, Name: %s\n", category.ID, category.Name)
	}
}

// GetCategoryByID ищет категорию по заданному ID и возвращает ее, если она найдена
func (cl *CategoryList) GetCategoryByID(id int64) (*Category, error) {
	for _, category := range cl.categories {
		if category.ID == id {
			return category, nil
		}
	}
	return nil, errors.New("category not found")
}

// FOR EXAMPLE
func GetCategoryListExample() *CategoryList {
	// Создаем новый пустой список категорий
	cl := newCategoryList()

	// Добавляем категории в список
	cl.addCategory(1, "Личные вещи")
	cl.addCategory(2, "Электроника")
	cl.addCategory(3, "Дом и сад")
	// и т.д. добавляем другие категории по мере необходимости

	return cl
}
