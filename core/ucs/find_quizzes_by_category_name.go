package ucs

import (
	"github.com/rcmendes/learnify/gameplay/core/entities"
	"github.com/rcmendes/learnify/gameplay/core/ucs/ports"
)

type findQuizByCategoryName struct {
	repo ports.QuizRepository
}

//MakeFindQuizByCategoryName creates a List All Quizzes Use Case instance.
func MakeFindQuizByCategoryName(
	repo ports.QuizRepository,
) ports.FindQuizByCategoryName {

	return &findQuizByCategoryName{
		repo,
	}
}

//TODO Handle errors or create custom ones.

func (uc *findQuizByCategoryName) FindByCategoryName(categoryName string) (entities.QuizList, error) {
	list, err := uc.repo.FindQuizByCategoryName(categoryName)

	if err != nil {
		return nil, err
	}

	return list, nil
}
