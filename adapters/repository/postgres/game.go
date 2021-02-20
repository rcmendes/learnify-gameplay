package postgres

import (
	"github.com/go-pg/pg/v10"
	"github.com/rcmendes/learnify/gameplay/adapters/repository/postgres/models"
	"github.com/rcmendes/learnify/gameplay/core/entities"
	"github.com/rcmendes/learnify/gameplay/core/ucs/ports"
)

type gamePostgresRepository struct {
	connectFunc func() *pg.DB
}

func NewGamePostgresRepository() ports.GameRepository {
	return &gamePostgresRepository{
		connectFunc: Connect,
	}
}

func (repo *gamePostgresRepository) Insert(game entities.Game) error {
	conn := repo.connectFunc()
	defer conn.Close()

	model := models.GameModel{}
	model.Load(game)

	_, err := conn.Conn().Model(&model).Insert()
	if err != nil {
		//TODO Handle error
		return err
	}

	for _, quiz := range model.Quizzes {
		quiz.GameID = model.ID
		_, err := conn.Model(quiz).Insert()
		if err != nil {
			//TODO Handle error
			return err
		}
	}

	return nil
}

func (repo *gamePostgresRepository) Update(game entities.Game, propagate bool) error {
	conn := repo.connectFunc()
	defer conn.Close()

	model := models.GameModel{}
	model.Load(game)

	_, err := conn.Conn().Model(&model).UpdateNotZero()
	if err != nil {
		//TODO Handle error
		return err
	}

	if propagate {
		for _, quiz := range model.Quizzes {
			quiz.GameID = model.ID
			_, err := conn.Model(quiz).UpdateNotZero()
			if err != nil {
				//TODO Handle error
				return err
			}
		}
	}
	return nil
}

func (repo *gamePostgresRepository) GetByID(id entities.GameID, lazy bool) (*entities.Game, error) {
	//TODO Handle Lazy and Handle GameQuiz
	conn := repo.connectFunc()
	defer conn.Close()

	var model models.GameModel

	if err := conn.Conn().Model(&model).Where("id=?", id).First(); err != nil {
		//TODO Handle error
		return nil, err
	}

	game := model.To()
	return &game, nil
}
