package ucs

import (
	"fmt"

	"github.com/rcmendes/learnify/gameplay/core/entities"
	"github.com/rcmendes/learnify/gameplay/core/ucs/ports"
)

type validateAnswerGameQuiz struct {
	repo ports.GameRepository
}

func MakeValidateAnswerGameQuiz(repo ports.GameRepository) ports.ValidateAnswerGameQuiz {
	return &validateAnswerGameQuiz{
		repo,
	}
}

// EvaluateQuizAnswer evaluates if the answer for the quiz is correct.
func (uc *validateAnswerGameQuiz) ValidateAnswer(gameID entities.GameID, quizID entities.QuizID, answer entities.QuizID) (bool, error) {
	//TODO Add game as parameter and create an external ID for each option of the quiz (QameQuizID, maybe base 64 of gameID+quizID?)
	game, err := uc.repo.GetByID(gameID, false)

	if err != nil {
		//TODO Handle error
		return false, err
	}

	gameQuiz := game.GetQuizByID(quizID)

	if gameQuiz == nil {
		//TODO Wrap error
		return false, fmt.Errorf("No quiz with ID '%s' was found in the Game '%s'", quizID, gameID)
	}

	result := gameID == gameQuiz.Quiz.ID

	if result {
		gameQuiz.Status = entities.GameQuizStatus.Correct
	} else {
		gameQuiz.Status = entities.GameQuizStatus.Wrong

	}

	err = uc.repo.Update(*game, true)
	if err != nil {
		//TODO Handle error
		return false, err
	}

	return result, nil
}
