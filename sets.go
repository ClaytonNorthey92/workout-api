package main

import (
	"errors"

	"github.com/ClaytonNorthey92/workout-api/models"
	"github.com/gobuffalo/uuid"
)

type SetParams struct {
	RoutineID uuid.UUID
	UserID    uuid.UUID
}

func PerformSet(s SetParams) error {
	set := models.Set{
		UserID:    s.UserID,
		RoutineID: s.RoutineID,
	}

	validateErrors, err := set.Save(db)
	if err != nil {
		return err
	}

	if validateErrors.Count() > 0 {
		return errors.New(validateErrors.Error())
	}

	return nil
}
