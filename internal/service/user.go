package service

import (
	"fmt"
	"context"
	"github.com/rs/zerolog"
	"github.com/zakkbob/mxguard/internal/model"
)

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
	user, err := u.Repo.CreateUser(ctx, params)
	if err != nil {
		return user, fmt.Errorf("creating user: %w", err)
	}
	return user, nil
}
