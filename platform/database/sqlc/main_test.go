package database

import (
	"database/sql"
	"errors"
	env "github.com/fabiotavarespr/go-env"
	logger "github.com/fabiotavarespr/go-logger"
	"github.com/fabiotavarespr/go-logger/attributes"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries

const (
	dbDriver = "postgres"
	dbSource = "postgres://root:simple-bank@localhost:5432/simple-bank?sslmode=disable"
)

func TestMain(m *testing.M) {
	logger.Init(&logger.Option{
		ServiceName:    env.GetEnvWithDefaultAsString("SERVICE", "Testing"),
		ServiceVersion: env.GetEnvWithDefaultAsString("VERSION", "0.0.1"),
		Environment:    env.GetEnvWithDefaultAsString("ENVIRONMENT", "dev"),
		LogLevel:       env.GetEnvWithDefaultAsString("LOG_LEVEL", "debug"),
	})

	defer logger.Sync()

	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		details := attributes.New().WithError(errors.New("interrupt signal detected"))
		details["dbDriver"] = dbDriver
		details["dbSource"] = dbSource

		logger.Fatal("Connection fail", details)
	}

	testQueries = New(conn)
	os.Exit(m.Run())
}
