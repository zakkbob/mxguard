package modeltestutils

// Used for generating random data for tests
// ONLY for tests

import (
	"testing"

	"github.com/brianvoe/gofakeit/v7"

	"github.com/google/uuid"
	"github.com/zakkbob/mxguard/internal/model"
)

// Returns a user with completely random fields
// testing.T is passed so that errors can be automatically checked
func RandUser(t *testing.T) model.User {
	id, err := uuid.NewRandom()
	if err != nil {
		t.Errorf("failed to generate random UUID: %v", err)
	}

	user, err := model.MakeUser(id, gofakeit.LetterN(20), gofakeit.Email(), gofakeit.Bool())
	if err != nil {
		t.Errorf("failed to make random user: %v", err)
	}
	return user
}

// Returns a user with completely random fields
// testing.T is passed so that errors can be automatically checked
func RandAlias(t *testing.T) model.Alias {
	alias, err := model.MakeAlias(gofakeit.LetterN(20), gofakeit.Sentence(10))
	if err != nil {
		t.Errorf("failed to make random alias: %v", err)
	}
	return alias
}
