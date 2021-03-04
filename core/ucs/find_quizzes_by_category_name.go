package ucs

import (
	"github.com/rcmendes/learnify-gameplay/core/entities"
	"github.com/rcmendes/learnify-gameplay/core/ucs/ports"
)

type findQuizByCategoryName struct {
	categoryRepo ports.CategoryRepository
	quizRepo     ports.QuizRepository
}

//MakeFindQuizByCategoryName creates a List All Quizzes Use Case instance.
func MakeFindQuizByCategoryName(
	categoryRepo ports.CategoryRepository,
	quizRepo ports.QuizRepository,

) ports.FindQuizByCategoryName {

	return &findQuizByCategoryName{
		categoryRepo,
		quizRepo,
	}
}

//TODO Handle errors or create custom ones.

func (uc *findQuizByCategoryName) FindByCategoryName(category string) (entities.QuizList, error) {
	cat, err := uc.categoryRepo.GetByName(category)

	if err != nil {
		//TODO handle error
		return nil, err
	}

	list, err := uc.quizRepo.FindQuizByCategoryID(cat.ID)

	if err != nil {
		//TODO handle error
		return nil, err
	}

	return list, nil
}
