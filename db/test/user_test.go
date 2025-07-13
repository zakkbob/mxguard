package db_test

// Tests for the UserRepository
// A database mock is not used, because this will couple the tests too closely to the database schema, meaning they will be too brittle
// Instead, dockertest is used to spin up a real database, which can then be treated like a black box, where the tests only verify the output of functions (model.User and errors)

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"
	"github.com/zakkbob/mxguard/db"
	"github.com/zakkbob/mxguard/internal/model"
	"github.com/zakkbob/mxguard/internal/service"
	"testing"
)

func TestCreatingAValidUserThenGettingByIDReturnsNoErrors(t *testing.T) {
	userRepo := db.NewPostgresUserRepository(dbPool)

	username := "TCAVUTGBIRNO" // Initialisation of test name
	isAdmin := false

	createdUser, err := userRepo.CreateUser(context.Background(), service.CreateUserParams{
		Username: username,
		IsAdmin:  isAdmin,
	})

	assert.NoError(t, err, "CreateUser should not return an error")
	assert.Equal(t, username, createdUser.Username, "Created user's username should match requested username")
	assert.Equal(t, isAdmin, createdUser.IsAdmin, "Created user's admin status should match requested admin status")

	retrievedUser, err := userRepo.GetUserByID(context.Background(), createdUser.ID)

	assert.NoError(t, err, "GetUserByID should not return an error")
	assert.Equal(t, createdUser, retrievedUser, "Retrieved user should match created user")

}

func TestGettingNonExistentUserByIDReturnsErrUserNotFound(t *testing.T) {
	userRepo := db.NewPostgresUserRepository(dbPool)

	id, _ := uuid.FromBytes([]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})

	retrievedUser, err := userRepo.GetUserByID(context.Background(), id)

	assert.ErrorIs(t, err, service.ErrUserNotFound, "GetUserByID should return ErrUserNotFound")
	assert.Equal(t, model.User{}, retrievedUser, "Retrieved user should be empty")
}

func TestErrInternalReturnedUponInternalError(t *testing.T) {
	brokenPool, err := pgxpool.New(context.Background(), "")
	userRepo := db.NewPostgresUserRepository(brokenPool)

	id, _ := uuid.FromBytes([]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})

	retrievedUser, err := userRepo.GetUserByID(context.Background(), id)

	var ei *service.ErrInternal
	assert.ErrorAs(t, err, &ei, "PostgresRepository should return ErrInternal")
	assert.Equal(t, model.User{}, retrievedUser, "Retrieved user should be empty")
}
