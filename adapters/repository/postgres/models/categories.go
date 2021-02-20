package models

import (
	"fmt"

	"github.com/rcmendes/learnify/gameplay/core/entities"
)

type CategoryID = storableID

//CategoryModel defines the properties of a CategoryModel database entity.
type CategoryModel struct {
	Storable
	tableName   struct{} `pg:"categories"`
	Name        string   `pg:"name"`
	Description *string  `pg:"description"`
}

func (c *CategoryModel) String() string {
	return fmt.Sprintf("User<name=%s, description=%q, storable=(%s)>",
		c.Name, *c.Description, c.Storable.String())
}

func (c *CategoryModel) To() entities.Category {
	return entities.Category{
		ID:          c.ID,
		Name:        c.Name,
		Description: c.Description,
	}
}

func NewCategoryModel(category entities.Category) CategoryModel {
	return CategoryModel{
		Storable:    Storable{ID: category.ID},
		Name:        category.Name,
		Description: category.Description,
	}
}
