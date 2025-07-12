package database

import (
	"context"
	"fmt"

	"github.com/rs/zerolog"
	"github.com/zakkbob/mxguard/internal/config"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Conn interface {
	Begin(ctx context.Context) (pgx.Tx, error)
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, optionsAndArgs ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, optionsAndArgs ...interface{}) pgx.Row
}

func Init(logger zerolog.Logger, cfg *config.Config) *pgxpool.Pool {
	var pool *pgxpool.Pool
	var ctx = context.Background()

	// Initialise the connection pool
	var err error
	var url = fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=%s", cfg.Postgres.User, cfg.Postgres.Password, cfg.Postgres.Url, cfg.Postgres.DB, cfg.Postgres.SSLmode)
	pool, err = pgxpool.New(ctx, url)
	if err != nil {
		logger.Fatal().Err(err).Str("url", url).Msg("Unable to connect to database")
	}

	// Verify the connection
	if err = pool.Ping(ctx); err != nil {
		logger.Fatal().Err(err).Str("url", url).Msg("Unable to ping database")
	}

	logger.Info().Str("url", url).Msg("Connected to PostgreSQL database")

	return pool
}
