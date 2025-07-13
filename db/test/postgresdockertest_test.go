package db_test

import (
	"context"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"

	_ "github.com/lib/pq"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/rs/zerolog/log"
	"github.com/jackc/pgx/v5/pgxpool"
)

var dbPool *pgxpool.Pool

func TestMain(m *testing.M) {
	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	dockerPool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatal().Msgf("Could not construct dockertest pool: %s", err)
	}

	err = dockerPool.Client.Ping()
	if err != nil {
		log.Fatal().Msgf("Could not connect to Docker: %s", err)
	}

	// pulls an image, creates a container based on it and runs it
	resource, err := dockerPool.RunWithOptions(&dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "17",
		Env: []string{
			"POSTGRES_PASSWORD=secret",
			"POSTGRES_USER=user_name",
			"POSTGRES_DB=dbname",
			"listen_addresses = '*'",
		},
	}, func(config *docker.HostConfig) {
		// set AutoRemove to true so that stopped container goes away by itself
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
	})
	if err != nil {
		log.Fatal().Msgf("Could not start resource: %s", err)
	}

	hostAndPort := resource.GetHostPort("5432/tcp")
	databaseUrl := fmt.Sprintf("postgres://user_name:secret@%s/dbname?sslmode=disable", hostAndPort)

	log.Info().Msgf("Connecting to database on url: %s", databaseUrl)

	resource.Expire(120) // Tell docker to hard kill the container in 120 seconds

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	dockerPool.MaxWait = 120 * time.Second
	if err = dockerPool.Retry(func() error {
		dbPool, err = pgxpool.New(context.Background(), databaseUrl)
		if err != nil {
			return err 
		}

		return dbPool.Ping(context.Background())
	}); err != nil {
		log.Fatal().Msgf("Could not connect to database: %s", err)
	}

	wd, _ := os.Getwd()
	log.Info().Msgf("Working directory: %s", wd)

	migrater, err := migrate.New("file://../migrations", databaseUrl) //Should this be hardcoded? I guess I'll find out
	if err != nil {
		log.Fatal().Err(err).Msg("golang-migrate could not connect to database")
	}

	for err := migrater.Up(); err != nil && err != migrate.ErrNoChange; {
		log.Fatal().Err(err).Msg("Could not run migrations")
	}

	defer func() {
		if err := dockerPool.Purge(resource); err != nil {
			log.Fatal().Msgf("Could not purge resource: %s", err)
		}
	}()

	// run tests
	m.Run()
}

// Is my dockertest thing even working?
func TestDockerTest(t *testing.T) {
	err := dbPool.Ping(context.Background())
	assert.NoError(t, err, "Should be able to ping database")
}
