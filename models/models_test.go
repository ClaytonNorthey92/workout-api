package models

import (
	"os"
	"testing"

	"github.com/gobuffalo/pop"
)

var db *pop.Connection

func TestMain(m *testing.M) {

	var err error

	db, err = pop.Connect("test")
	if err != nil {
		panic(err.Error())
	}

	code := m.Run()

	users := []User{}

	db.All(&users)

	for _, user := range users {
		db.Destroy(&user)
	}

	db.Close()

	os.Exit(code)
}
