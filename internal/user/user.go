package user

import (
	"context"
	"fmt"
	"strconv"

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
        INSERT INTO user (username, is_admin)
        VALUES ($1, $2)
        RETURNING id
    `

	var ctx = context.Background()

	var id int
	err := conn.QueryRow(ctx, sql, username, isAdmin, false).Scan(&id)
	if err != nil {
		return fmt.Errorf("error creating task: %w", err)
	}

	log.Trace("Created user with ID: " + strconv.Itoa(id))
	return nil
}
