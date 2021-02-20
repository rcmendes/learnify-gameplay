package ucs

import (
	"github.com/rcmendes/learnify/gameplay/core/entities"
	"github.com/rcmendes/learnify/gameplay/core/ucs/ports"
)

type findAllCategories struct {
	repo ports.CategoryRepository
}

func MakeFindAllCategories(repo ports.CategoryRepository) ports.FindAllCategories {
	return &findAllCategories{
		repo,
	}
}

func (uc *findAllCategories) FindAll() (entities.CategoryList, error) {
	list, err := uc.repo.ListAll()

	if err != nil {
		return nil, err
	}

	return list, nil
}
