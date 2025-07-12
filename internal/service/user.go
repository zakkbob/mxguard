package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/rs/zerolog"
	"github.com/zakkbob/mxguard/internal/model"
)

var ErrEmptyUsername = errors.New("username cannot be empty")


type CreateUserParams struct {
	Username string
	IsAdmin  bool
}

type UserRepository interface {
	CreateUser(context.Context, CreateUserParams) (model.User, error)
}

func NewUserService(logger zerolog.Logger, repo UserRepository) *UserService {
	return &UserService{
		Logger: logger,
		Repo:   repo,
	}
}

type UserService struct {
	Logger zerolog.Logger
	Repo   UserRepository
}

func (u *UserService) CreateUser(ctx context.Context, params CreateUserParams) (model.User, error) {
	if params.Username == "" {
		return model.User{}, ErrEmptyUsername
	}

	user, err := u.Repo.CreateUser(ctx, params)
	if err != nil {
		return user, fmt.Errorf("creating user: %w", err)
	}
	return user, nil
}
