package ports

import "github.com/rcmendes/learnify-gameplay/core/entities"

type PlayerRepository interface {
	GetByID(id entities.PlayerID) (*entities.Player, error)
}
