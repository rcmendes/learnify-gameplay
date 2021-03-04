package ports

import "github.com/rcmendes/learnify-gameplay/core/entities"

//QuizRepository defines the contract of a Quiz entity repository.
type QuizRepository interface {
	ListAll() (entities.QuizList, error)
	FindQuizByCategoryID(categoryID entities.CategoryID) (entities.QuizList, error)
	GetQuizByID(id entities.QuizID) (*entities.Quiz, error)
	FindQuizzesSameCategory(id entities.QuizID) (entities.QuizList, error)
}

type FindAllQuizzes interface {
	FindAll() (entities.QuizList, error)
}

type FindQuizByCategoryName interface {
	FindByCategoryName(category string) (entities.QuizList, error)
}

type FindQuizByCategoryID interface {
	FindByCategoryID(categoryID entities.CategoryID) (entities.QuizList, error)
}

type FindQuizzesSameCategory interface {
	FindAllSameCategory(id entities.QuizID) (entities.QuizList, error)
}

type FindQuiz interface {
	FindByID(id entities.QuizID) (*entities.Quiz, error)
	GetImageByID(id entities.QuizID) (*entities.MediaInfo, error)
	GetAudioByID(id entities.QuizID) (*entities.MediaInfo, error)
}
