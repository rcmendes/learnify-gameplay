package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type storableID = uuid.UUID

//Storable defines some attributes of a database entity.
type Storable struct {
	ID        storableID `pg:"id"`
	CreatedAt *time.Time `pg:"created_at"`
	UpdatedAt *time.Time `pg:"updated_at"`
}

func (s *Storable) String() string {
	return fmt.Sprintf("Storable <ID=%d, createdAt=%q, updatedAt=%q>",
		s.ID, s.CreatedAt, s.UpdatedAt)
}
