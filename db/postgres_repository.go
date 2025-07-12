package db

import (
	"context"
	"github.com/google/uuid"
	"github.com/zakkbob/mxguard/internal/database"
	"github.com/zakkbob/mxguard/internal/model"
	"github.com/zakkbob/mxguard/internal/service"
	"fmt"
)

type PostgresUserRepository struct {
	Conn database.Conn 
}

var _ service.UserRepository = &PostgresUserRepository{}

func NewPostgresUserRepository(conn database.Conn) *PostgresUserRepository {
	return &PostgresUserRepository { 
		Conn: conn,
	}
}

func (u *PostgresUserRepository) CreateUser(ctx context.Context, username string, isAdmin bool) (model.User, error) {
	sql := `
        INSERT INTO usr (username, is_admin)
        VALUES ($1, $2)
        RETURNING id
    `

	var id uuid.UUID
	err := u.Conn.QueryRow(ctx, sql, username, isAdmin).Scan(&id)
	if err != nil {
		return model.User{}, fmt.Errorf("querying database: %w", err)
	}

	return model.User {
		ID: id,
		Username: username,
		IsAdmin: isAdmin,
	}, nil
}
