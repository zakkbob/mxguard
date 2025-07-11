package user

import (
	"context"

	"github.com/google/uuid"

	log "github.com/sirupsen/logrus"
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
	err := conn.QueryRow(ctx, sql, username, isAdmin).Scan(&id)
	if err != nil {
		return err
	}

	log.Trace("Created user with ID: " + id.String())
	return nil
}
