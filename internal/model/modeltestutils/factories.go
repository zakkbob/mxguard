package modeltestutils

// Used for generating random data for tests
// ONLY for tests

import (
	"math/rand"
	"github.com/google/uuid"
	"github.com/zakkbob/mxguard/internal/model"
	"testing"
)

func randStr(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	b := make([]byte, n)
	for i := 0; i < n; i++ {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}

func randBool() bool {
	return rand.Intn(2) == 0
}

// Returns a user with completely random fields
// testing.T is passed so that errors can be automatically checked
func RandUser(t *testing.T) model.User {
	id, err := uuid.NewRandom()
	if err != nil {
		t.Errorf("failed to generate random UUID: %v", err)
	}

	return model.User {
		ID: id,
		Username: randStr(20),
		IsAdmin: randBool(),
	}
}
