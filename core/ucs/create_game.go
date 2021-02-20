package ucs

import (
	"fmt"
	"math/rand"

	"github.com/google/uuid"
	"github.com/rcmendes/learnify/gameplay/core/entities"
	"github.com/rcmendes/learnify/gameplay/core/ucs/ports"
)

type createGame struct {
	gameRepo ports.GameRepository
	quizRepo ports.QuizRepository
}

func MakeCreateGame(gameRepo ports.GameRepository, quizRepo ports.QuizRepository) ports.CreateGame {
	return &createGame{
		gameRepo,
		quizRepo,
	}
}

//Create is a method that creates a game.
func (uc *createGame) Create(newGame entities.NewGameData) (*entities.GameID, error) {
	//TODO Validate if the player exists

	gameID := uuid.New()

	game := entities.NewGame(gameID, entities.NewPlayer(newGame.PlayerID))

	//TODO Replace for an application parameter.
	totalOfGameQuizzes := 4
	category := entities.Category{ID: newGame.CategoryID}
	quizzes, err := uc.quizRepo.FindQuizByCategoryName(category.Name)
	if err != nil {
		//TODO handle error
		return nil, err
	}

	length := len(quizzes)
	if length == 0 {
		//TODO create an error for this
		return nil, fmt.Errorf("no quiz found for this category")
	}

	for i := 0; i < length && i < totalOfGameQuizzes; i++ {
		index := rand.Intn(length - 1)
		quiz := *quizzes[index]
		if !game.Contains(quiz.ID) {
			game.AddQuiz(quiz)
		}
	}

	if err := uc.gameRepo.Insert(game); err != nil {
		//TODO handle error
		return nil, err
	}

	return &gameID, nil
}
