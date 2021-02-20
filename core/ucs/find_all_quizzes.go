package ucs

import (
	"github.com/rcmendes/learnify/gameplay/core/entities"
	"github.com/rcmendes/learnify/gameplay/core/ucs/ports"
)

type findAllQuizzes struct {
	repo ports.QuizRepository
}

//MakeFindAllQuizzes creates a List All Quizzes Use Case instance.
func MakeFindAllQuizzes(
	repo ports.QuizRepository,
) ports.FindAllQuizzes {

	return &findAllQuizzes{
		repo,
	}
}

//TODO Handle errors or create custom ones.

func (uc *findAllQuizzes) FindAll() (entities.QuizList, error) {
	list, err := uc.repo.ListAll()

	if err != nil {
		return nil, err
	}

	return list, nil
}
