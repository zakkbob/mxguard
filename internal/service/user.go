package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/zakkbob/mxguard/internal/model"
)

//var ErrNoID = errors.New("ID cannot be nil")
var ErrEmptyUsername = errors.New("username cannot be empty")


type CreateUserParams struct {
	Username string
	IsAdmin  bool
}

type UserRepository interface {
	CreateUser(context.Context, CreateUserParams) (model.User, error)
	DeleteUser(context.Context, model.User) error
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{
		Repo:   repo,
	}
}

type UserService struct {
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


func (u *UserService) DeleteUser(ctx context.Context, user model.User) error {
	//if user.ID == nil {
	//	return ErrNoID
	//}

	err := u.Repo.DeleteUser(ctx, user)
	if err != nil {
		return fmt.Errorf("deleting user: %w", err)
	}
	return nil
}
