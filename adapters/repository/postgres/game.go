package postgres

import (
	"github.com/go-pg/pg/v10"
	"github.com/rcmendes/learnify-gameplay/adapters/repository/postgres/models"
	"github.com/rcmendes/learnify-gameplay/core/entities"
	"github.com/rcmendes/learnify-gameplay/core/ucs/ports"
	"go.uber.org/zap"
)

type gamePostgresRepository struct {
	connectFunc func() *pg.DB
	logger      *zap.Logger
}

func NewGamePostgresRepository() ports.GameRepository {
	//TODO review this logger init (put as parameter and create an adapter)
	logger, _ := zap.NewDevelopment()
	return &gamePostgresRepository{
		connectFunc: Connect,
		logger:      logger,
	}
}

func (repo *gamePostgresRepository) Insert(game entities.Game) error {
	defer repo.logger.Sync()
	conn := connect()

	model := models.GameModel{}
	model.Load(game)

	query := `INSERT INTO games (id, player_id, status) VALUES (:id, :player_id, :status)`
	_, err := conn.NamedExec(query, model)
	if err != nil {
		//TODO HandleError
		// repo.logger.Error(err.String())
		return err
	}

	query = `INSERT INTO game_quizzes (game_id, quiz_id, status) VALUES (:game_id, :quiz_id, :status)`
	for _, quiz := range model.Quizzes {
		quiz.GameID = model.ID
		_, err := conn.NamedExec(query, quiz)
		if err != nil {
			//TODO HandleError
			// repo.logger.Error(err.String())
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

	game := model.ToEntity()
	return &game, nil
}
