package ucs

import (
	"github.com/google/uuid"
	"github.com/rcmendes/learnify/gameplay/core/entities"
	"github.com/rcmendes/learnify/gameplay/core/ucs/ports"
)

type createCategory struct {
	repo ports.CategoryRepository
}

func MakeCreateCategory(repo ports.CategoryRepository) ports.CreateCategory {
	return &createCategory{
		repo,
	}
}

func (uc *createCategory) Create(newCategory entities.NewCategory) (*entities.CategoryID, error) {
	id := uuid.New()
	category := newCategory.To(id)

	if err := uc.repo.Insert(category); err != nil {
		return nil, err
	}

	return &id, nil
}
