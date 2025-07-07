package database

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/zakkbob/mxguard/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

var pool *pgxpool.Pool
var ctx = context.Background()

type User struct {
	ID       int
	IsAdmin  bool
	Username string
}

func Init(c *config.Config) {
	// Initialise the connection pool
	var err error
	var url = fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=%s", c.Postgres.User, c.Postgres.Password, c.Postgres.Url, c.Postgres.DB, c.Postgres.SSLmode)
	pool, err = pgxpool.New(ctx, url)
	if err != nil {
		log.WithError(err).WithField("url", url).Fatal("Unable to connect to database")
	}

	// Verify the connection
	if err = pool.Ping(ctx); err != nil {
		log.WithError(err).WithField("url", url).Fatal("Unable to ping database")
	}

	log.WithField("url", url).Info("Connected to PostgreSQL database")
}
