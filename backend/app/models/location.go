package models

// Location представляет структуру для хранения информации о категории
type Location struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Parent   *Location
	Children []*Location
}

func emptyLocation(id int64, name string, parent *Location) Location {
	return Location{
		ID:       id,
		Name:     name,
		Parent:   parent,
		Children: nil,
	}
}

// Creates a new category and adds it to given node.
// Returns a newly created node
func (c *Location) addChild(id int64, name string) *Location {
	newLocation := Location{
		ID:       id,
		Name:     name,
		Parent:   c,
		Children: nil,
	}
	c.Children = append(c.Children, &newLocation)
	return &newLocation
}

func (c *Location) Traverse() []*Location {
	if len(c.Children) == 0 {
		return []*Location{c}
	}
	var result []*Location
	for _, child := range c.Children {
		result = append(result, child.Traverse()...)
	}
	result = append(result, c)
	return result
}

func GetLocationTreeExample() *Location {
	root := emptyLocation(1, "ROOT", nil)
	ivan := root.addChild(2, "Ивановская область")
	ivan.addChild(3, "Кинешма")
	ivan.addChild(4, "Заволжск")
	ivan.addChild(5, "Родники")

	return &root
}
