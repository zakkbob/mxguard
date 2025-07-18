package service_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/zakkbob/mxguard/internal/model"
	"github.com/zakkbob/mxguard/internal/model/modeltestutils"
	"github.com/zakkbob/mxguard/internal/service"
)

// Captures the passed parameters
type MockUserRepository struct {
	User             model.User  // canned user value
	Alias            model.Alias // canned alias value
	DeletedUsername  string
	DeletedID        uuid.UUID
	DeletedUser      model.User
	CreateUserParams service.CreateUserParams
	UserDeleted      bool
	UserCreated      bool
	AliasCreated     bool
}

func (m *MockUserRepository) GetUserByID(ctx context.Context, id uuid.UUID) (model.User, error) {
	return m.User, nil
}

func (m *MockUserRepository) GetUserByUsername(ctx context.Context, username string) (model.User, error) {
	return m.User, nil
}

func (m *MockUserRepository) DeleteUserByID(ctx context.Context, id uuid.UUID) error {
	m.UserDeleted = true
	m.DeletedID = id
	return nil
}

func (m *MockUserRepository) DeleteUserByUsername(ctx context.Context, username string) error {
	m.UserDeleted = true
	m.DeletedUsername = username
	return nil
}

func (m *MockUserRepository) CreateUser(ctx context.Context, params service.CreateUserParams) (model.User, error) {
	m.UserCreated = true
	m.CreateUserParams = params
	return m.User, nil
}

func (m *MockUserRepository) CreateAlias(ctx context.Context, user model.User, name string, description string) (model.Alias, error) {
	m.AliasCreated = true
	return m.Alias, nil
}

func TestCreatingAliasDoesntModifyAlias(t *testing.T) {
	expected := modeltestutils.RandAlias(t)

	userRepo := &MockUserRepository{Alias: expected}
	userService := service.NewUserService(userRepo)

	user := modeltestutils.RandUser(t)
	got, err := userService.CreateAlias(context.Background(), user, expected.Name(), expected.Description())

	assert.NoError(t, err, "CreateAlias should not return an error")
	assert.True(t, userRepo.AliasCreated)

	modeltestutils.AssertAliasesEqual(t, expected, got, "Returned alias should be equal to created alias")
}

func TestCreatingUserDoesntModifyUser(t *testing.T) {
	expected := modeltestutils.RandUser(t)
	params := service.CreateUserParams{
		Username: expected.Username(),
		IsAdmin:  expected.IsAdmin(),
		Email:    expected.Email(),
	}

	userRepo := &MockUserRepository{User: expected}
	userService := service.NewUserService(userRepo)

	got, err := userService.CreateUser(context.Background(), params)

	assert.NoError(t, err, "CreateUser should not return an error")
	assert.Equal(t, params, userRepo.CreateUserParams, "Passed user params should be equal to requested")

	modeltestutils.AssertUsersEqual(t, expected, got, "Returned user should be equal to created user")
}

func TestCreatingUserWithEmptyUsernameThrowsErrEmptyUsername(t *testing.T) {
	params := service.CreateUserParams{
		Username: "",
		IsAdmin:  false,
		Email:    "test@email.com",
	}
	expected := model.User{}
	userRepo := &MockUserRepository{}
	userService := service.NewUserService(userRepo)

	got, err := userService.CreateUser(context.Background(), params)

	assert.ErrorIs(t, err, service.ErrEmptyUsername, "CreateUser should return ErrEmptyUsername")
	assert.Equal(t, expected, got, "Returned user should be empty")
	assert.False(t, userRepo.UserCreated, "No user should've been created")
	assert.False(t, userRepo.UserDeleted, "No user should've been deleted")
}

func TestCreatingUserWithEmptyEmailThrowsErrEmptyEmail(t *testing.T) {
	params := service.CreateUserParams{
		Username: "test",
		IsAdmin:  false,
		Email:    "",
	}
	expected := model.User{}
	userRepo := &MockUserRepository{}
	userService := service.NewUserService(userRepo)

	got, err := userService.CreateUser(context.Background(), params)

	assert.ErrorIs(t, err, service.ErrEmptyEmail, "CreateUser should return ErrEmptyEmail")
	assert.Equal(t, expected, got, "Returned user should be empty")
	assert.False(t, userRepo.UserCreated, "No user should've been created")
	assert.False(t, userRepo.UserDeleted, "No user should've been deleted")
}

func TestDeleteUser(t *testing.T) {
	user := modeltestutils.RandUser(t)

	userRepo := &MockUserRepository{}
	userService := service.NewUserService(userRepo)

	err := userService.DeleteUser(context.Background(), user)

	assert.NoError(t, err, "DeleteUser should not return an error")
	assert.False(t, userRepo.UserCreated, "No user should've been created")
	assert.True(t, userRepo.UserDeleted, "A user should've been deleted")
	assert.Equal(t, user.ID(), userRepo.DeletedID, "Deleted user's ID should equal requested user's ID")
}
