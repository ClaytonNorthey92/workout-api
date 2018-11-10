package main

import (
	"os"
	"testing"

	"github.com/ClaytonNorthey92/workout-api/models"
	"github.com/gobuffalo/pop"
)

var db *pop.Connection

func TestMain(m *testing.M) {

	var err error

	db, err = pop.Connect("test")
	if err != nil {
		panic(err.Error())
	}

	sets := []models.Set{}
	db.All(&sets)
	for _, set := range sets {
		err := db.Destroy(&set)
		if err != nil {
			panic(err.Error())
		}
	}

	users := []models.User{}

	db.All(&users)

	for _, user := range users {
		err := db.Destroy(&user)
		if err != nil {
			panic(err.Error())
		}
	}

	code := m.Run()

	db.Close()

	os.Exit(code)
}
