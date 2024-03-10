package models

// Category представляет структуру для хранения информации о категории
type Category struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Parent   *Category
	Children []*Category
}

func emptyCategory(id int64, name string, parent *Category) Category {
	return Category{
		ID:       id,
		Name:     name,
		Parent:   parent,
		Children: nil,
	}
}

// Creates a new category and adds it to given node.
// Returns a newly created node
func (c *Category) addChild(id int64, name string) *Category {
	newCategory := Category{
		ID:       id,
		Name:     name,
		Parent:   c,
		Children: nil,
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

// FOR EXAMPLE
func GetCategoryListExample() []*Category {
	// Создаем новый пустой список категорий
	root := emptyCategory(1, "ROOT", nil)
	electronics := root.addChild(2, "Бытовая электроника")
	electronics.addChild(3, "Товары для компьютера")
	electronics.addChild(4, "Фототехника")
	electronics.addChild(8, "Ноутбуки")

	return root.traverse()
}