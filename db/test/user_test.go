package db_test

// Tests for the UserRepository
// A database mock is not used, because this will couple the tests too closely to the database schema, meaning they will be too brittle
// Instead, dockertest is used to spin up a real database, which can then be treated like a black box, where the tests only verify the output of functions (model.User and errors)

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"
	"github.com/zakkbob/mxguard/db"
	"github.com/zakkbob/mxguard/internal/model"
	"github.com/zakkbob/mxguard/internal/model/modeltestutils"
	"github.com/zakkbob/mxguard/internal/service"
)

// CreateUser isn't the SUT, so an error will just fail the test and log, not be returned
func RandUserCreate(t *testing.T, r service.UserRepository) model.User {
	// the ID is ignored here, but this saves creating a specific CreateUserParams random generator. Do i even need that struct?
	randUser := modeltestutils.RandUser(t)

	createdUser, err := r.CreateUser(context.Background(), service.CreateUserParams{
		Username: randUser.Username(),
		IsAdmin:  randUser.IsAdmin(),
		Email:    randUser.Email(),
	})
	if err != nil {
		t.Errorf("failed to create user: %v", err)
	}

	return createdUser
}

func TestCreateAliasReturnsNoError(t *testing.T) {
	userRepo := db.NewPostgresUserRepository(dbPool)

	user := modeltestutils.RandUser(t)
	alias := modeltestutils.RandAlias(t)

	createdUser, err := userRepo.CreateUser(context.Background(), service.CreateUserParams{
		Username: user.Username(),
		IsAdmin:  user.IsAdmin(),
		Email:    user.Email(),
	})

	assert.NoError(t, err, "CreateUser should not return an error")

	_, err = userRepo.CreateAlias(context.Background(), createdUser, alias.Name(), alias.Description())
	assert.NoError(t, err, "CreateAlias should not return an error")
}

func TestCreateAndGetUserByID(t *testing.T) {
	userRepo := db.NewPostgresUserRepository(dbPool)

	user := modeltestutils.RandUser(t)

	createdUser, err := userRepo.CreateUser(context.Background(), service.CreateUserParams{
		Username: user.Username(),
		IsAdmin:  user.IsAdmin(),
		Email:    user.Email(),
	})

	assert.NoError(t, err, "CreateUser should not return an error")
	modeltestutils.AssertUsersEqualIgnoreID(t, user, createdUser, "Created user should match requested user")

	retrievedUser, err := userRepo.GetUserByID(context.Background(), createdUser.ID())

	assert.NoError(t, err, "GetUserByID should not return an error")
	assert.Equal(t, createdUser, retrievedUser, "Retrieved user should match created user")
}

func TestCreateAndDeleteUserByID(t *testing.T) {
	userRepo := db.NewPostgresUserRepository(dbPool)
	user := RandUserCreate(t, userRepo)

	err := userRepo.DeleteUserByID(context.Background(), user.ID())
	assert.NoError(t, err, "DeleteUserByID should not return an error")

	_, err = userRepo.GetUserByID(context.Background(), user.ID())
	assert.ErrorIs(t, err, service.ErrUserNotFound, "User should be deleted")

	err = userRepo.DeleteUserByID(context.Background(), user.ID())
	assert.ErrorIs(t, err, service.ErrUserNotFound, "DeleteUserByID should return ErrUserNotFound")
}

func TestCreateAndDeleteUserByUsername(t *testing.T) {
	userRepo := db.NewPostgresUserRepository(dbPool)
	user := RandUserCreate(t, userRepo)

	err := userRepo.DeleteUserByUsername(context.Background(), user.Username())
	assert.NoError(t, err, "DeleteUser should not return an error")

	_, err = userRepo.GetUserByID(context.Background(), user.ID())
	assert.ErrorIs(t, err, service.ErrUserNotFound, "User should be deleted")

	err = userRepo.DeleteUserByUsername(context.Background(), user.Username())
	assert.ErrorIs(t, err, service.ErrUserNotFound, "DeleteUserByUsername should return ErrUserNotFound")
}

func TestGetUserByIDWhenUserDoesntExistReturnsErrUserNotFound(t *testing.T) {
	userRepo := db.NewPostgresUserRepository(dbPool)

	id, _ := uuid.FromBytes([]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})

	retrievedUser, err := userRepo.GetUserByID(context.Background(), id)

	assert.ErrorIs(t, err, service.ErrUserNotFound, "GetUserByID should return ErrUserNotFound")
	assert.Equal(t, model.User{}, retrievedUser, "Retrieved user should be empty")
}

func TestGetUserByIDInternalErrorReturnsErrInternal(t *testing.T) {
	brokenPool, err := pgxpool.New(context.Background(), "")
	if err != nil {
		t.Errorf("failed to construct broken pool: %v", err)
	}

	userRepo := db.NewPostgresUserRepository(brokenPool)

	id, _ := uuid.FromBytes([]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})

	retrievedUser, err := userRepo.GetUserByID(context.Background(), id)

	var ei *service.ErrInternal
	assert.ErrorAs(t, err, &ei, "PostgresRepository should return ErrInternal")
	assert.Equal(t, model.User{}, retrievedUser, "Retrieved user should be empty")
}
