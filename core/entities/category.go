package entities

import "github.com/google/uuid"

type NewCategory struct {
	Name        string
	Description *string
}

type CategoryID = uuid.UUID

type Category struct {
	ID          CategoryID
	Name        string
	Description *string
}

type CategoryList = []*Category

func (c *NewCategory) To(id CategoryID) Category {
	return Category{
		ID:          id,
		Name:        c.Name,
		Description: c.Description,
	}
}
