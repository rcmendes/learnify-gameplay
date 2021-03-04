package postgres

import (
	"github.com/rcmendes/learnify-gameplay/adapters/repository/postgres/models"
	"github.com/rcmendes/learnify-gameplay/core/entities"
	"github.com/rcmendes/learnify-gameplay/core/ucs/ports"
)

type playerPostgresRepository struct{}

func NewPlayerPostgresRepository() ports.PlayerRepository {
	return &playerPostgresRepository{}
}

func (repo *playerPostgresRepository) GetByID(id entities.PlayerID) (*entities.Player, error) {
	conn := connect()

	model := models.PlayerModel{}

	if err := conn.Get(&model, "SELECT * FROM players WHERE id=$1", id); err != nil {
		//TODO HandleError
		return nil, err
	}

	player := model.ToEntity()

	return player, nil
}
