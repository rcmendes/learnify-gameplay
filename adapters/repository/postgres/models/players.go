package models

import (
	"github.com/google/uuid"
	"github.com/rcmendes/learnify-gameplay/core/entities"
)

type PlayerModel struct {
	ID   uuid.UUID
	Name string
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
