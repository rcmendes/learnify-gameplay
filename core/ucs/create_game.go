package ucs

import (
	"fmt"
	"math/rand"

	"github.com/google/uuid"
	"github.com/rcmendes/learnify-gameplay/core/entities"
	"github.com/rcmendes/learnify-gameplay/core/ucs/ports"
)

type createGame struct {
	gameRepo   ports.GameRepository
	quizRepo   ports.QuizRepository
	playerRepo ports.PlayerRepository
}

func MakeCreateGame(gameRepo ports.GameRepository, quizRepo ports.QuizRepository, playerRepo ports.PlayerRepository) ports.CreateGame {
	return &createGame{
		gameRepo,
		quizRepo,
		playerRepo,
	}
}

//Create is a method that creates a game.
func (uc *createGame) Create(newGame ports.CreateGameInput) (*entities.GameID, error) {
	playerID := newGame.PlayerID()

	player, err := uc.playerRepo.GetByID(playerID)
	if err != nil {
		//TODO Validate if the player exists
		//TODO Handle Error
		return nil, err
	}

	gameID := uuid.New()

	game := entities.NewGame(gameID, *player)

	//TODO Replace for an application parameter.
	totalOfGameQuizzes := 4
	// category := entities.Category{ID: newGame.CategoryID(), Name: "animais"}
	quizzes, err := uc.quizRepo.FindQuizByCategoryID(newGame.CategoryID())
	if err != nil {
		//TODO handle error
		return nil, err
	}

	length := len(quizzes)
	if length == 0 {
		//TODO create an error for this
		return nil, fmt.Errorf("no quiz found for this category")
	}

	//TODO use shuffle
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
