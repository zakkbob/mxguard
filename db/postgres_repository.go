package db

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/zakkbob/mxguard/internal/model"
	"github.com/zakkbob/mxguard/internal/service"
)

type PostgresUserRepository struct {
	Conn Conn
}

// Verify that PostgresUserRepository conforms to service.UserRepository
var _ service.UserRepository = &PostgresUserRepository{}

func NewPostgresUserRepository(conn Conn) *PostgresUserRepository {
	return &PostgresUserRepository{
		Conn: conn,
	}
}

type postgresUser struct {
	ID       uuid.UUID
	Username string
	IsAdmin  bool
	Email    string
}

func (u postgresUser) ConstructDomainUser() (model.User, error) {
	user, err := model.MakeUser(u.ID, u.Username, u.Email, u.IsAdmin)
	if err != nil {
		return user, fmt.Errorf("constructing user from database: %w", err)
	}
	return user, nil
}

func (u *PostgresUserRepository) CreateUser(ctx context.Context, params service.CreateUserParams) (model.User, error) {
	sql := `
        INSERT INTO usr (username, is_admin, email)
        VALUES ($1, $2, $3)
        RETURNING id
    `

	var id uuid.UUID
	err := u.Conn.QueryRow(ctx, sql, params.Username, params.IsAdmin, params.Email).Scan(&id)
	if err != nil {
		return model.User{}, fmt.Errorf("querying database: %w", &service.ErrInternal{Err: err})
	}

	user := postgresUser{
		ID:       id,
		Username: params.Username,
		IsAdmin:  params.IsAdmin,
		Email:    params.Email,
	}

	return user.ConstructDomainUser()
}

func (u *PostgresUserRepository) GetUserByID(ctx context.Context, id uuid.UUID) (model.User, error) {
	sql := `
	SELECT id, username, is_admin, email FROM usr
	WHERE id = $1
	`

	var user postgresUser
	err := u.Conn.QueryRow(ctx, sql, id).Scan(&user.ID, &user.Username, &user.IsAdmin, &user.Email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.User{}, service.ErrUserNotFound
		}
		return model.User{}, fmt.Errorf("querying database: %w", &service.ErrInternal{Err: err})
	}

	return user.ConstructDomainUser()
}

func (u *PostgresUserRepository) GetUserByUsername(ctx context.Context, username string) (model.User, error) {
	sql := `
	SELECT id, username, is_admin, email FROM usr
	WHERE username = $1
	`

	var user postgresUser
	err := u.Conn.QueryRow(ctx, sql, username).Scan(&user.ID, &user.Username, &user.IsAdmin, &user.Email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.User{}, service.ErrUserNotFound
		}
		return model.User{}, fmt.Errorf("querying database: %w", &service.ErrInternal{Err: err})
	}

	return user.ConstructDomainUser()
}

func (u *PostgresUserRepository) DeleteUserByID(ctx context.Context, id uuid.UUID) error {
	sql := `
	DELETE FROM usr
	WHERE id = $1
    `

	commandTag, err := u.Conn.Exec(ctx, sql, id)
	if err != nil {
		return fmt.Errorf("querying database: %w", &service.ErrInternal{Err: err})
	}

	if commandTag.RowsAffected() == 0 {
		return service.ErrUserNotFound
	}

	return nil
}

func (u *PostgresUserRepository) DeleteUserByUsername(ctx context.Context, username string) error {
	sql := `
	DELETE FROM usr
	WHERE username = $1
    `

	commandTag, err := u.Conn.Exec(ctx, sql, username)
	if err != nil {
		return fmt.Errorf("querying database: %w", &service.ErrInternal{Err: err})
	}

	if commandTag.RowsAffected() == 0 {
		return service.ErrUserNotFound
	}

	return nil
}
