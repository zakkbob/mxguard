package user

import (
	"context"

	"github.com/google/uuid"

	"github.com/zakkbob/mxguard/internal/database"
)

type User struct {
	ID       int
	IsAdmin  bool
	Username string
}

func CreateUser(conn database.Conn, username string, isAdmin bool) error {
	sql := `
        INSERT INTO usr (username, is_admin)
        VALUES ($1, $2)
        RETURNING id
    `

	var ctx = context.Background()

	var id uuid.UUID
	return conn.QueryRow(ctx, sql, username, isAdmin).Scan(&id)
}
