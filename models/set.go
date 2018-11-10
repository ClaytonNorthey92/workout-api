package models

import (
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
)

type Set struct {
	ID        uuid.UUID `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	UserID    uuid.UUID `json:"user_id" db:"user_id"`
	RoutineID uuid.UUID `json:"routine_id" db:"routine_id"`
}

func (s *Set) Save(tx *pop.Connection) (*validate.Errors, error) {
	return tx.ValidateAndCreate(s)
}
