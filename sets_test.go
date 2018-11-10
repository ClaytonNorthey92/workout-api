package main

import (
	"fmt"
	"testing"

	"github.com/ClaytonNorthey92/workout-api/models"
)

func Test_AssociateSetWithUserAndRoutine(t *testing.T) {
	user := models.User{
		Username: "test user",
	}

	validateErrors, err := user.Save(db)
	if validateErrors.Error() != "" {
		t.Error(validateErrors.Error())
	}

	if err != nil {
		t.Error(err.Error())
	}

	var routine models.Routine

	query := db.Where("name = ?", "barbell flat bench press")
	err = query.First(&routine)
	if err != nil {
		t.Error(err.Error())
	}

	err = PerformSet(SetParams{
		UserID:    user.ID,
		RoutineID: routine.ID,
	})

	if err != nil {
		t.Error(err.Error())
	}

	var foundSet models.Set
	err = db.Last(&foundSet)
	if err != nil {
		t.Error(err.Error())
	}

	if foundSet.RoutineID != routine.ID {
		t.Error(fmt.Sprintf("invalid routine id: %s != %s", foundSet.RoutineID, routine.ID))
	}

	if foundSet.UserID != user.ID {
		t.Error(fmt.Sprintf("invalid user id: %s != %s", foundSet.UserID, user.ID))
	}
}
