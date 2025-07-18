package model

import (
	"errors"

	"github.com/google/uuid"
)

// Alias struct should be immutable after creation to avoid bugs
// Also forces use of the constructor
type Alias struct {
	name        string
	description string
	enabled     bool
}

func (a Alias) Name() string {
	return a.name
}

func (a Alias) Description() string {
	return a.description
}

func (a Alias) Enabled() bool {
	return a.enabled
}

func MakeAlias(name string, description string) (Alias, error) {
	if name == "" {
		return Alias{}, errors.New("can't have unnamed alias")
	}

	return Alias{
		name:        name,
		description: description,
		enabled:     true,
	}, nil
}

// User struct should be immutable after creation to avoid bugs
// Also forces use of the constructor
type User struct {
	id       uuid.UUID
	username string
	isAdmin  bool
	email    string
	aliases  []Alias
}

func (u User) ID() uuid.UUID {
	return u.id
}

func (u User) Username() string {
	return u.username
}

func (u User) IsAdmin() bool {
	return u.isAdmin
}

func (u User) Email() string {
	return u.email
}

func (u User) Aliases() []Alias {
	return u.aliases
}

// Unmarshal created User struct from database data, should not be used elsewhere
func UnmarshalUser(id uuid.UUID, username string, email string, isAdmin bool, aliases []Alias) (User, error) {
	return User{
		id:       id,
		username: username,
		isAdmin:  isAdmin,
		email:    email,
		aliases:  aliases,
	}, nil
}

func MakeUser(id uuid.UUID, username string, email string, isAdmin bool) (User, error) {
	if username == "" {
		return User{}, errors.New("cannot have empty username")
	}

	if email == "" {
		return User{}, errors.New("cannot have empty email")
	}

	return User{
		id:       id,
		username: username,
		isAdmin:  isAdmin,
		email:    email,
		aliases:  []Alias{},
	}, nil
}
