package models

import (
	"fmt"

	"github.com/rcmendes/learnify/gameplay/core/entities"
)

type QuizID = storableID

//QuizModel defines the properties of a QuizModel database entity.
type QuizModel struct {
	Storable
	tableName     struct{} `pg:"quizzes"`
	ImageFilename string   `pg:"image_filename"`
	Category      string   `pg:"category"`
	Palavra       string   `pg:"palavra"`
	Mot           string   `pg:"mot"`
	AudioFilename string   `pg:"audio_filename"`
}

func (q *QuizModel) String() string {
	return fmt.Sprintf("Quiz <url=%s, category=%s, "+
		"palavra=%s, mot=%s, audio=%s, storable=(%s)>",
		q.ImageFilename, q.Category, q.Palavra, q.Mot, q.AudioFilename,
		q.Storable.String())
}

func NewQuizModel(quiz entities.Quiz) QuizModel {
	return QuizModel{
		Storable:      Storable{ID: quiz.ID},
		ImageFilename: quiz.Image.Name,
		AudioFilename: quiz.Audio.Name,
		Palavra:       quiz.Palavra,
		Mot:           quiz.Mot,
		Category:      quiz.Category,
	}
}

func (q *QuizModel) To() entities.Quiz {
	return entities.Quiz{
		ID:       q.ID,
		Image:    entities.MediaInfo{Name: q.ImageFilename},
		Audio:    entities.MediaInfo{Name: q.AudioFilename},
		Palavra:  q.Palavra,
		Mot:      q.Mot,
		Category: q.Category,
	}
}
