package models

import (
	"testing"
)

func TestCreateUser(t *testing.T) {
	u := User{
		Username: "Clayton",
	}

	errs, _ := u.Save(db)
	if errs.Error() != "" {
		t.Error(errs.Error())
	}
}

func TestShouldFailOnBlankUsername(t *testing.T) {
	u := User{}

	errs, _ := u.Save(db)
	if errs.Error() != "Username can not be blank." {
		t.Error("received incorrect error")
	}
}

func TestShouldNotAllowDuplicateUsernames(t *testing.T) {
	userOne := User{
		Username: "blah",
	}

	userTwo := User{
		Username: "blah",
	}

	_, err := userOne.Save(db)
	if err != nil {
		t.Error(err.Error())
	}

	_, err = userTwo.Save(db)
	if err.Error() != "pq: duplicate key value violates unique constraint \"unique_username\"" {
		t.Error(err.Error())
	}
}
