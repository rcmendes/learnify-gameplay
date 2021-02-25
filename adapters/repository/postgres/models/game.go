package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/rcmendes/learnify-gameplay/core/entities"
)

type GameModel struct {
	// Storable
	ID        storableID  `db:"id"`
	CreatedAt *time.Time  `db:"created_at"`
	UpdatedAt *time.Time  `db:"updated_at"`
	tableName struct{}    `db:"game"`
	PlayerID  uuid.UUID   `db:"player_id"`
	Status    uint8       `db:"status"`
	Quizzes   []*GameQuiz //`pg:"rel:has-many"`
}

func (m *GameModel) Load(game entities.Game) {
	m.ID = game.ID
	m.PlayerID = game.Player.ID
	m.Status = game.Status

	m.Quizzes = make([]*GameQuiz, 0, len(game.Quizzes))
	for _, q := range game.Quizzes {
		gq := GameQuiz{
			GameID: game.ID,
			QuizID: q.Quiz.ID,
			Status: q.Status,
		}
		m.Quizzes = append(m.Quizzes, &gq)
	}
}

func (m *GameModel) To() entities.Game {
	quizzes := make([]*entities.GameQuiz, 0, len(m.Quizzes))
	for _, q := range m.Quizzes {
		quiz := entities.Quiz{}
		quiz.ID = q.QuizID
		//TODO evaluate lazy loading for the remainder Quiz info
		gq := entities.GameQuiz{
			Quiz:   quiz,
			Status: q.Status,
		}
		quizzes = append(quizzes, &gq)
	}

	return entities.Game{
		ID:      m.ID,
		Status:  m.Status,
		Quizzes: quizzes,
	}
}

type GameQuiz struct {
	GameID uuid.UUID `db:"game_id"`
	QuizID uuid.UUID `db:"quiz_id"`
	Status uint8     `db:"status"`
}
