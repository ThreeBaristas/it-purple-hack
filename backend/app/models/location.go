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
	return c.FindAllByPredicate(func(c *Location) bool {
		return true
	})
}

type LocationPredicate func(l *Location) bool

func (l *Location) FindAllByPredicate(predicate LocationPredicate) []*Location {
	var result []*Location
	for _, child := range l.Children {
		result = append(result, child.FindAllByPredicate(predicate)...)
	}
	if predicate(l) {
		result = append(result, l)
	}
	return result
}

func GetLocationTreeExample() *Location {
	root := emptyLocation(1, "ROOT", nil)
	ivan := root.addChild(2, "Ивановская область")
	ivan.addChild(3, "Кинешма")
	ivan.addChild(4, "Заволжск")
	ivan.addChild(5, "Родники")
	spb := root.addChild(20, "Санкт_Петербург")
	spb.addChild(21, "Петроградский р-н")
	return &root
}
