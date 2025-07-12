package db 

import (
	"context"
	"fmt"

	"github.com/zakkbob/mxguard/internal/config"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Conn interface {
	Ping(ctx context.Context) error
	Begin(ctx context.Context) (pgx.Tx, error)
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, optionsAndArgs ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, optionsAndArgs ...interface{}) pgx.Row
}

func InitConn(cfg *config.Config) (*pgxpool.Pool, error) {
	var pool *pgxpool.Pool
	var ctx = context.Background()

	// Initialise the connection pool
	var err error
	var url = fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=%s", cfg.Postgres.User, cfg.Postgres.Password, cfg.Postgres.Url, cfg.Postgres.DB, cfg.Postgres.SSLmode)
	pool, err = pgxpool.New(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("initialising connection pool: %w")
	}

	// Verify the connection
	if err = pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("pinging database: %w")
	}

	return pool, nil
}
