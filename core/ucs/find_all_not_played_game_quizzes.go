package ucs

import (
	"math/rand"
	"time"

	"github.com/rcmendes/learnify/gameplay/core/entities"
	"github.com/rcmendes/learnify/gameplay/core/ucs/ports"
)

type findOneNotPlayedGameQuiz struct {
	repo ports.GameRepository
}

func MakeFindOneNotPlayedGameQuiz(repo ports.GameRepository) ports.FindOneNotPlayedGameQuiz {
	return &findOneNotPlayedGameQuiz{
		repo,
	}
}

// GetNotPlayedQuiz returns a Quiz of Game that was not played yet. Returns nil if all
func (uc *findOneNotPlayedGameQuiz) FindOneNotPlayedQuiz(
	id entities.GameID) (*entities.GameQuiz, error) {
	//TODO Add game as parameter and create an external ID for each option of the quiz (QameQuizID, maybe base 64 of gameID+quizID?)
	game, err := uc.repo.GetByID(id, false)

	if err != nil {
		//TODO Handle error
		return nil, err
	}

	notPlayedList := game.GetNotPlayedQuizzes()

	shuffleQuizzes(notPlayedList)

	var gq *entities.GameQuiz
	if len(notPlayedList) > 0 {
		gq = notPlayedList[0]
	}

	return gq, nil
}

func shuffleQuizzes(quizzes []*entities.GameQuiz) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(quizzes), func(i, j int) { quizzes[i], quizzes[j] = quizzes[j], quizzes[i] })
}
