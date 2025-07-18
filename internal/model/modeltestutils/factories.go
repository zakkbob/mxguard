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

	return model.User{
		ID:       id,
		Username: gofakeit.LetterN(20),
		IsAdmin:  gofakeit.Bool(),
		Email:    gofakeit.Email(),
	}
}
