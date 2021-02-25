package ports

import (
	"github.com/google/uuid"
	"github.com/rcmendes/learnify-gameplay/core/entities"
)

type GameRepository interface {
	Insert(game entities.Game) error
	GetByID(id entities.GameID, lazy bool) (*entities.Game, error)
	Update(game entities.Game, propagate bool) error
}

//CreateGameInput defines the data to create a new Game.
type CreateGameInput interface {
	PlayerID() uuid.UUID
	CategoryID() uuid.UUID
}

type CreateGame interface {
	Create(input CreateGameInput) (*entities.GameID, error)
}

type ValidateAnswerGameQuiz interface {
	ValidateAnswer(
		gameID entities.GameID,
		quizID entities.QuizID,
		answer entities.QuizID) (bool, error)
}

type FindOneNotPlayedGameQuiz interface {
	FindOneNotPlayedQuiz(id entities.GameID) (*entities.GameQuiz, error)
}
