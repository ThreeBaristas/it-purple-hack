package models

import (
	"fmt"

	"go.uber.org/zap"
)

// Category представляет структуру для хранения информации о категории
type Category struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	DistToRoot int
	Parent     *Category
	Children   []*Category
}

func emptyCategory(id int64, name string, parent *Category) Category {
	category := Category{
		ID:       id,
		Name:     name,
		Parent:   parent,
		Children: nil,
	}
	if parent != nil {
		category.DistToRoot = parent.DistToRoot + 1
	}
	return category
}

// Creates a new category and adds it to given node.
// Returns a newly created node
func (c *Category) addChild(id int64, name string) *Category {
	newCategory := Category{
		ID:         id,
		Name:       name,
		Parent:     c,
		DistToRoot: c.DistToRoot + 1,
		Children:   nil,
	}
	c.Children = append(c.Children, &newCategory)
	return &newCategory
}

func (c *Category) traverse() []*Category {
	if len(c.Children) == 0 {
		return []*Category{c}
	}
	var result []*Category
	for _, child := range c.Children {
		result = append(result, child.traverse()...)
	}
	result = append(result, c)
	return result
}

func (c *Category) FindChildById(id int64) *Category {
	if c.ID == id {
		return c
	}
	for _, child := range c.Children {
		res := child.FindChildById(id)
		if res != nil {
			return res
		}
	}
	return nil
}

type CategoryPredicate func(c *Category) bool

func (c *Category) FindAllByPredicate(predicate CategoryPredicate) []*Category {
	var result []*Category
	for _, child := range c.Children {
		result = append(result, child.FindAllByPredicate(predicate)...)
	}
	if predicate(c) {
		result = append(result, c)
	}
	return result
}

// FOR EXAMPLE
func GetCategoryTreeExample() *Category {
	// Создаем новый пустой список категорий
	// root := emptyCategory(1, "ROOT", nil)
	// electronics := root.addChild(2, "Бытовая электроника")
	// electronics.addChild(3, "Товары для компьютера")
	// electronics.addChild(4, "Фототехника")
	// electronics.addChild(8, "Ноутбуки")
	// vehicles := root.addChild(20, "Транспорт")
	// vehicles.addChild(21, "Машины")
	// vehicles.addChild(22, "Катера")
	// vehicles.addChild(23, "Самолеты")

	return GetCategoryTreeHuge()
}

func GetCategoryTreeHuge() *Category {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	logger.Info("Started generating categories tree")
	const h = 4
	const nodesOnLevel = 10
	var id int64 = 2
	// Total length of 10^4 = 10_000

	var generate func(c *Category, hCur int)
	generate = func(c *Category, hCur int) {
		if hCur == 0 {
			return
		}
		for i := 0; i < nodesOnLevel; i++ {
			c.addChild(id, fmt.Sprintf("Category #%d", id))
			id++
			generate(c, hCur-1)
		}
	}

	// Total of 10_000 nodes
	root := emptyCategory(1, "ROOT", nil)
	logger.Info("Started generating categories tree")
	generate(&root, h)
	logger.Info("Generated categoires tree. Traversing it to find len")
	arr := root.traverse()
	logger.Info("Traversed categories tree", zap.Int("len", len(arr)))
	return &root
}
