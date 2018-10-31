package models

import (
	"testing"
)

func TestCreateUser(t *testing.T) {
	u := User{
		Username: "Clayton",
	}

	errs := u.Save(db)
	if errs.Error() != "" {
		t.Error(errs.Error())
	}
}

func TestShouldFailOnBlankUsername(t *testing.T) {
	u := User{}

	errs := u.Save(db)
	if errs.Error() != "Username can not be blank." {
		t.Error("received incorrect error")
	}
}
