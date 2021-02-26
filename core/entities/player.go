package entities

import "github.com/google/uuid"

//PlayerID defines the type of the ID of a Player
type PlayerID = uuid.UUID

type Player struct {
	ID   PlayerID
	Name string
}

func NewPlayer(id PlayerID, name string) Player {
	return Player{
		ID:   id,
		Name: name,
	}
}
