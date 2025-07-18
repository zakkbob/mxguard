package model

import (
	"errors"

	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID
	Username string
	IsAdmin  bool
	Email    string
}

func NewUser(id uuid.UUID, username string, isAdmin bool, email string) (User, error) {
	if username == "" {
		return User{}, errors.New("cannot have empty username")
	}

	if email == "" {
		return User{}, errors.New("cannot have empty email")
	}

	return User{
		ID:       id,
		Username: username,
		IsAdmin:  isAdmin,
		Email:    email,
	}, nil
}
