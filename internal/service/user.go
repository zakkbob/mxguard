package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/zakkbob/mxguard/internal/model"
)

// ---- Errors ----

// var ErrNoID = errors.New("ID cannot be nil")
var (
	ErrEmptyUsername = errors.New("username cannot be empty")
	ErrUserNotFound = errors.New("user not found in repository")
)

// Represents an internal repository error
// Abtracts the specific error, while making it accessible through .Error()
// Consumers can check if the error is internal using Error.Is and handle appropriately
type ErrInternal struct {
	Err error
}

func (e *ErrInternal) Error() string {
	return fmt.Sprintf("internal repository error: %v", e.Err)
}

// ----------------

type CreateUserParams struct {
	Username string
	IsAdmin  bool
}

type UserRepository interface {
	CreateUser(context.Context, CreateUserParams) (model.User, error)
	DeleteUser(context.Context, model.User) error
	GetUserByID(context.Context, uuid.UUID) (model.User, error)
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{
		Repo: repo,
	}
}

type UserService struct {
	Repo UserRepository
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
