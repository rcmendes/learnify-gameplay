package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/rcmendes/learnify-gameplay/core/entities"
)

type PlayerModel struct {
	CreatedAt *time.Time `db:"created_at"`
	ID        uuid.UUID  `db:"id"`
	Name      string     `db:"name"`
}

func NewPlayerModel(player entities.Player) *PlayerModel {
	return &PlayerModel{
		ID:   player.ID,
		Name: player.Name,
	}
}

func (p *PlayerModel) ToEntity() *entities.Player {
	return &entities.Player{
		ID:   p.ID,
		Name: p.Name,
	}
}
