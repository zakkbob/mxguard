package service_test

import (
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/zakkbob/mxguard/internal/model"
	"github.com/zakkbob/mxguard/internal/model/modeltestutils"
	"github.com/zakkbob/mxguard/internal/service"
	"testing"
)

// Captures the passed parameters
type MockUserRepository struct {
	DeletedUsername  string
	DeletedID        uuid.UUID
	DeletedUser      model.User
	CreateUserParams service.CreateUserParams
	UserDeleted      bool
	UserCreated      bool
}

func (m *MockUserRepository) GetUserByID(ctx context.Context, id uuid.UUID) (model.User, error) {
	return model.User{}, nil
}

func (m *MockUserRepository) GetUserByUsername(ctx context.Context, username string) (model.User, error) {
	return model.User{}, nil
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
	return model.User{
		Username: params.Username,
		IsAdmin:  params.IsAdmin,
	}, nil
}

func TestCreatingUserDoesntModifyUser(t *testing.T) {
	params := service.CreateUserParams{
		Username: "success",
		IsAdmin:  false,
	}
	expected := model.User{
		Username: "success",
		IsAdmin:  false,
	}
	userRepo := &MockUserRepository{}
	userService := service.NewUserService(userRepo)

	got, err := userService.CreateUser(context.Background(), params)

	assert.NoError(t, err, "CreateUser should not return an error")
	assert.Equal(t, params, userRepo.CreateUserParams, "Passed user params should be equal to requested")
	assert.Equal(t, expected, got, "Returned user should be equal to created user")
}

func TestCreatingUserWithEmptyUsernameThrowsErrEmptyUsername(t *testing.T) {
	params := service.CreateUserParams{
		Username: "",
		IsAdmin:  false,
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

func TestDeleteUser(t *testing.T) {
	user := modeltestutils.RandUser(t) 

	userRepo := &MockUserRepository{}
	userService := service.NewUserService(userRepo)

	err := userService.DeleteUser(context.Background(), user)

	assert.NoError(t, err, "DeleteUser should not return an error")
	assert.False(t, userRepo.UserCreated, "No user should've been created")
	assert.True(t, userRepo.UserDeleted, "A user should've been deleted")
	assert.Equal(t, user.ID, userRepo.DeletedID, "Deleted user's ID should equal requested user's ID")
}
