package user

import (
	"context"
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
        INSERT INTO usr (username, is_admin)
        VALUES ($1, $2)
        RETURNING id
    `

	var ctx = context.Background()

	var id string
	err := conn.QueryRow(ctx, sql, username, isAdmin).Scan(&id)
	if err != nil {
		log.WithError(err).Error("Failed to create user")
		return err
	}

	log.Trace("Created user with ID: " + strconv.Itoa(id))
	return nil
}
