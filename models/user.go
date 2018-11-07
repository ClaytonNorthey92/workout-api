package models

import (
	"fmt"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

type User struct {
	ID        uuid.UUID `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Username  string    `json:"username" db:"username"`
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (u *User) Validate(tx *pop.Connection) (*validate.Errors, error) {
	userCount, err := u.ExistingUsernameCount(tx)
	if err != nil {
		return nil, err
	}

	return validate.Validate(
		&validators.StringIsPresent{Field: u.Username, Name: "Username"},
		&validators.IntIsLessThan{Name: "count", Field: userCount, Compared: 1, Message: fmt.Sprintf("duplicate username: \"%s\"", u.Username)},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (u *User) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (u *User) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// Save is the method used to persist user data (save a user) to the database
func (u *User) Save(tx *pop.Connection) (*validate.Errors, error) {
	return tx.ValidateAndCreate(u)
}

// ExistingUsernameCount will return the count of users with the same username or an error
// if one has occurred
func (u *User) ExistingUsernameCount(tx *pop.Connection) (int, error) {
	where := tx.Where("username = ?", u.Username)
	return where.Count(&User{})
}
