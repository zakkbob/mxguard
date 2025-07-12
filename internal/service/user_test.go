package service_test

import (
	"context"
	"io"
	"testing"
	"github.com/rs/zerolog"
	"github.com/zakkbob/mxguard/internal/service"
	"github.com/stretchr/testify/assert"
	"github.com/zakkbob/mxguard/internal/model"
)

type MockUserRepository struct{
}

func (m *MockUserRepository) CreateUser(ctx context.Context, params service.CreateUserParams) (model.User, error) {
	return model.User {
		Username: params.Username,
		IsAdmin: params.IsAdmin,
	}, nil
}

func TestValidUserSucceeds(t *testing.T) {
	params := service.CreateUserParams {
		Username: "success",
		IsAdmin: false,
	}
	expected := model.User{
		Username: "success",
		IsAdmin: false,
	}
	userService := service.NewUserService(
		zerolog.New(io.Discard), 
		&MockUserRepository{},
	)

	got, err := userService.CreateUser(context.TODO(), params)
	
	assert.NoError(t, err, "should not error")
	assert.Equal(t, expected, got, "should be equal")
}

func TestEmptyUsernameThrowsErrow(t *testing.T) {
	params := service.CreateUserParams {
		Username: "",
		IsAdmin: false,
	}
	expected := model.User{}
	userService := service.NewUserService(
		zerolog.New(io.Discard), 
		&MockUserRepository{},
	)

	got, err := userService.CreateUser(context.TODO(), params)
	
	assert.ErrorIs(t, err, service.ErrEmptyUsername, "should not error")
	assert.Equal(t, expected, got, "should be equal")
}
