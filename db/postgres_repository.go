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

type postgresAlias struct {
	Name        string
	Description string
	Enabled     bool
}

type postgresUser struct {
	ID       uuid.UUID
	Username string
	IsAdmin  bool
	Email    string
	Aliases  []postgresAlias
}

func (u postgresUser) constructDomainUser() model.User {
	aliases := make([]model.Alias, len(u.Aliases))

	for i, alias := range u.Aliases {
		aliases[i] = alias.constructDomainAlias()
	}

	user := model.UnmarshalUser(u.ID, u.Username, u.Email, u.IsAdmin, aliases)
	return user
}

func (a postgresAlias) constructDomainAlias() model.Alias {
	return model.UnmarshalAlias(a.Name, a.Description, a.Enabled)
}

func (u *PostgresUserRepository) CreateAlias(ctx context.Context, user model.User, name string, description string) (model.Alias, error) {
	sql := `
        INSERT INTO alias (usr_id, name, description, enabled)
        VALUES ($1, $2, $3, $4)
    `

	_, err := u.Conn.Exec(ctx, sql, user.ID(), name, description, true)
	if err != nil {
		return model.Alias{}, fmt.Errorf("querying database: %w", &service.ErrInternal{Err: err})
	}

	alias := postgresAlias{
		Name:        name,
		Description: description,
		Enabled:     true,
	}

	return alias.constructDomainAlias(), nil
}

func (u *PostgresUserRepository) getUserAliases(ctx context.Context, userID uuid.UUID) ([]postgresAlias, error) {
	sql := `
	SELECT name, description, enabled FROM alias
	WHERE usr_id = $1
	`

	rows, err := u.Conn.Query(ctx, sql, userID)
	if err != nil {
		return []postgresAlias{}, err
	}

	defer rows.Close()

	var aliases []postgresAlias
	var alias postgresAlias
	for rows.Next() {
		err = rows.Scan(&alias.Name, &alias.Description, &alias.Enabled)
		if err != nil {
			return []postgresAlias{}, err
		}
		aliases = append(aliases, alias)
	}

	if rows.Err() != nil {
		return []postgresAlias{}, rows.Err()
	}

	return aliases, nil

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

	return user.constructDomainUser(), nil
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
		return model.User{}, fmt.Errorf("querying database for user: %w", &service.ErrInternal{Err: err})
	}

	user.Aliases, err = u.getUserAliases(ctx, id)
	if err != nil {
		return model.User{}, fmt.Errorf("querying database for user aliases: %w", &service.ErrInternal{Err: err})
	}

	return user.constructDomainUser(), nil
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
		return model.User{}, fmt.Errorf("querying database for user: %w", &service.ErrInternal{Err: err})
	}

	user.Aliases, err = u.getUserAliases(ctx, user.ID)
	if err != nil {
		return model.User{}, fmt.Errorf("querying database for user aliases: %w", &service.ErrInternal{Err: err})
	}

	return user.constructDomainUser(), nil
}

func (u *PostgresUserRepository) DeleteUserByID(ctx context.Context, id uuid.UUID) error {
	sql := `
	DELETE FROM usr
	WHERE id = $1
    `

	commandTag, err := u.Conn.Exec(ctx, sql, id)
	if err != nil {
		return fmt.Errorf("deleting user from database: %w", &service.ErrInternal{Err: err})
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
		return fmt.Errorf("deleting user from database: %w", &service.ErrInternal{Err: err})
	}

	if commandTag.RowsAffected() == 0 {
		return service.ErrUserNotFound
	}

	return nil
}
