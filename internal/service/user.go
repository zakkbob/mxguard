package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/zakkbob/mxguard/internal/model"
)

// ---- Errors ----
// TODO: May be too granular right now, I will check over them once the minimum demo is working

var (
	ErrEmptyUsername    = errors.New("username cannot be empty")
	ErrEmptyEmail       = errors.New("email cannot be empty")
	ErrEmptyName        = errors.New("name cannot be empty")
	ErrEmptyDescription = errors.New("description cannot be empty")
	ErrUserNotFound     = errors.New("user not found")
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

// TODO: Is this needed? may remove later
type CreateUserParams struct {
	Username string
	IsAdmin  bool
	Email    string
}

type UserRepository interface {
	CreateUser(context.Context, CreateUserParams) (model.User, error)
	DeleteUserByID(context.Context, uuid.UUID) error
	DeleteUserByUsername(context.Context, string) error
	GetUserByID(context.Context, uuid.UUID) (model.User, error)
	GetUserByUsername(context.Context, string) (model.User, error)

	CreateAlias(ctx context.Context, user model.User, name string, description string) (model.Alias, error)
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

	if params.Email == "" {
		return model.User{}, ErrEmptyEmail
	}

	user, err := u.Repo.CreateUser(ctx, params)
	if err != nil {
		return user, fmt.Errorf("creating user: %w", err)
	}
	return user, nil
}

func (u *UserService) GetUserByID(ctx context.Context, id uuid.UUID) (model.User, error) {
	user, err := u.Repo.GetUserByID(ctx, id)
	if err != nil {
		return model.User{}, fmt.Errorf("getting user: %w", err)
	}
	return user, nil
}

func (u *UserService) GetUserByUsername(ctx context.Context, username string) (model.User, error) {
	user, err := u.Repo.GetUserByUsername(ctx, username)
	if err != nil {
		return model.User{}, fmt.Errorf("getting user: %w", err)
	}
	return user, nil
}

// Shorthand for DeleteUserByID(ctx, user.ID)
func (u *UserService) DeleteUser(ctx context.Context, user model.User) error {
	return u.DeleteUserByID(ctx, user.ID())
}

func (u *UserService) DeleteUserByID(ctx context.Context, id uuid.UUID) error {
	err := u.Repo.DeleteUserByID(ctx, id)
	if err != nil {
		return fmt.Errorf("deleting user: %w", err)
	}
	return nil
}

func (u *UserService) DeleteUserByUsername(ctx context.Context, username string) error {
	err := u.Repo.DeleteUserByUsername(ctx, username)
	if err != nil {
		return fmt.Errorf("deleting user: %w", err)
	}
	return nil
}

func (u *UserService) CreateAlias(ctx context.Context, user model.User, name string, description string) (model.Alias, error) {
	if name == "" {
		return model.Alias{}, ErrEmptyName
	}

	if name == "" {
		return model.Alias{}, ErrEmptyDescription
	}

	alias, err := u.Repo.CreateAlias(ctx, user, name, description)
	if err != nil {
		return model.Alias{}, fmt.Errorf("creating alias: %w", err)
	}
	return alias, err
}
