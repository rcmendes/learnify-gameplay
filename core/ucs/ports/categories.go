package ports

import "github.com/rcmendes/learnify/gameplay/core/entities"

//CategoryRepository defines the services that a Cetegory repository must provide.
type CategoryRepository interface {
	Insert(category entities.Category) error
	ListAll() (entities.CategoryList, error)
}

type CreateCategory interface {
	Create(newCategory entities.NewCategory) (*entities.CategoryID, error)
}

type FindAllCategories interface {
	FindAll() (entities.CategoryList, error)
}
