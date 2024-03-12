package models

import (
	"fmt"

	"go.uber.org/zap"
)

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
	return GetLocationTreeHuge()
}

func GetLocationTreeHuge() *Location {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	logger.Info("Started generating locations tree")
	const h = 5
	const nodesOnLevel = 7
	var id int64 = 2
	// Total length of ~ 7^5 ~ 1608

	var generate func(c *Location, hCur int)
	generate = func(c *Location, hCur int) {
		if hCur == 0 {
			return
		}
		for i := 0; i < nodesOnLevel; i++ {
			c.addChild(id, fmt.Sprintf("Location #%d", id))
			id++
			generate(c, hCur-1)
		}
	}

	// Total of 10_000 nodes
	root := emptyLocation(1, "ROOT", nil)
	logger.Info("Started generating locations tree")
	generate(&root, h)
	logger.Info("Generated locations tree. Traversing it to find len")
	arr := root.Traverse()
	logger.Info("Traversed locations tree", zap.Int("len", len(arr)))
	return &root
}
