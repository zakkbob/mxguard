package db_test

// Tests for the UserRepository
// A database mock is not used, because this will couple the tests too closely to the database schema, meaning they will be too brittle
// Instead, dockertest is used to spin up a real database, which can then be treated like a black box, where the tests only verify the output of functions (model.User and errors)

import (
	"testing"
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/zakkbob/mxguard/db"
	"github.com/zakkbob/mxguard/internal/service"
)

func TestCreatingAValidUserThenGettingByIDReturnsNoErrors(t *testing.T) {
	userRepo := db.NewPostgresUserRepository(conn)

	username := "TCAVUTGBIRNO" // Initialisation of test name 
	isAdmin := false
	
	createdUser, err := userRepo.CreateUser(context.Background(), service.CreateUserParams{
		Username: username,
		IsAdmin: isAdmin,
	})

	assert.NoError(t, err, "CreateUser should not return an error")
	assert.Equal(t, username, createdUser.Username, "Created user's username should match requested username")
	assert.Equal(t, isAdmin, createdUser.IsAdmin, "Created user's admin status should match requested admin status")


	retrievedUser, err := userRepo.GetUserByID(context.Background(), createdUser.ID)

	assert.NoError(t, err, "GetUserByID should not return an error")
	assert.Equal(t, createdUser, retrievedUser, "Retrieved user should match created user")

}
