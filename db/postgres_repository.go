package db

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/zakkbob/mxguard/internal/database"
	"github.com/zakkbob/mxguard/internal/model"
	"github.com/zakkbob/mxguard/internal/service"
)

type PostgresUserRepository struct {
	Conn database.Conn
}

var _ service.UserRepository = &PostgresUserRepository{}

func NewPostgresUserRepository(conn database.Conn) *PostgresUserRepository {
	return &PostgresUserRepository{
		Conn: conn,
	}
}

func (u *PostgresUserRepository) CreateUser(ctx context.Context, params service.CreateUserParams) (model.User, error) {
	sql := `
        INSERT INTO usr (username, is_admin)
        VALUES ($1, $2)
        RETURNING id
    `

	var id uuid.UUID
	err := u.Conn.QueryRow(ctx, sql, params.Username, params.IsAdmin).Scan(&id)
	if err != nil {
		return model.User{}, fmt.Errorf("querying database: %w", err)
	}

	return model.User{
		ID:       id,
		Username: params.Username,
		IsAdmin:  params.IsAdmin,
	}, nil
}
