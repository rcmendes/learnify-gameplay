package entities

import (
	"github.com/google/uuid"
)

//GameID defines the type of the ID of a Game
type GameID = uuid.UUID

//NewGameData defines the data to create a new Game.
type NewGameData struct {
	PlayerID   PlayerID
	CategoryID CategoryID
}

//Game defines the structure of a Game.
type Game struct {
	ID      GameID
	Player  Player
	Status  statusType
	Quizzes []*GameQuiz
}

//GameQuiz defines the structure of Quizzes associated to a Game.
type GameQuiz struct {
	Quiz   Quiz
	Status statusType
}

type statusType = uint8

type gameStatusOptions struct {
	//Created represents the status of a game when it's created.
	Created statusType
	//Playing represents the status of a game when it's been played.
	Playing statusType
	//Finished represents the status of a game when it's finished.
	Finished statusType
}

// GameStatus defines the available status of a Game.
var GameStatus = gameStatusOptions{
	Created:  1,
	Playing:  2,
	Finished: 3,
}

type gameQuizStatusOptions struct {
	//NotPlayed represents the status of a Quiz when it was not yet been played.
	NotPlayed statusType
	//Correct represents the status of a Quiz was its answer was correc.
	Correct statusType
	//Wrong represents the status of a Quiz when its answer was wrong.
	Wrong statusType
}

// GameQuizStatus defines the available status of a GameQuiz.
var GameQuizStatus = gameQuizStatusOptions{
	NotPlayed: 1,
	Correct:   2,
	Wrong:     3,
}

func NewGame(id GameID, player Player) Game {
	return Game{
		ID:     id,
		Player: player,
		Status: GameStatus.Created,
	}
}

func NewGameQuiz(quiz Quiz) GameQuiz {
	return GameQuiz{
		Quiz:   quiz,
		Status: GameQuizStatus.NotPlayed,
	}
}

// GetQuizByID returns the corresponding GameQuiz, if it exists, by using the Quiz ID.
func (g *Game) GetQuizByID(id QuizID) *GameQuiz {
	for _, quiz := range g.Quizzes {
		if quiz.Quiz.ID == id {
			return quiz
		}
	}

	return nil
}

// GetNotPlayedQuizzes returns the Quizzes of a Game that were not yet been played.
func (g *Game) GetNotPlayedQuizzes() []*GameQuiz {
	var notPlayedQuizzes []*GameQuiz
	for _, quiz := range g.Quizzes {
		if quiz.Status == GameQuizStatus.NotPlayed {
			notPlayedQuizzes = append(notPlayedQuizzes, quiz)
		}
	}

	return notPlayedQuizzes
}

func (g *Game) AddQuiz(quiz Quiz) {
	gq := NewGameQuiz(quiz)
	g.Quizzes = append(g.Quizzes, &gq)
}

func (g *Game) Contains(id QuizID) bool {
	for _, gq := range g.Quizzes {
		if gq.Quiz.ID == id {
			return true
		}
	}

	return false
}
