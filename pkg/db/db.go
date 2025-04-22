package db

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pgx-contrib/pgxotel"
)

// Useful for mocking dbpools
type dbconn interface {
	Query(ctx context.Context, sql string, optionsAndArgs ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, optionsAndArgs ...interface{}) pgx.Row
	Close()
}

var Pool dbconn

func Init() error {
	config, err := Config()

	if err != nil {
		return err
	}

	connPool, err := pgxpool.NewWithConfig(context.Background(), config)

	if err != nil {
		return err
	}

	connection, err := connPool.Acquire(context.Background())

	if err != nil {
		return err
	}

	defer connection.Release()

	err = connection.Ping(context.Background())

	if err != nil {
		return fmt.Errorf("could not reach the database: %w", err)
	}

	Pool = connPool

	fmt.Println("Connected to the database")

	return nil
}

func Config() (*pgxpool.Config, error) {
	const defaultMaxConns = int32(4)
	const defaultMinConns = int32(0)
	const defaultMaxConnLifetime = time.Hour
	const defaultMaxConnIdleTime = 30 * time.Minute
	const defaultHealthCheckPeriod = time.Minute
	const defaultConnectTimeout = 5 * time.Second

	dbConfig, err := pgxpool.ParseConfig(os.Getenv("DATABASE_URL"))

	if err != nil {
		return nil, err
	}

	dbConfig.MaxConns = defaultMaxConns
	dbConfig.MinConns = defaultMinConns
	dbConfig.MaxConnLifetime = defaultMaxConnLifetime
	dbConfig.MaxConnIdleTime = defaultMaxConnIdleTime
	dbConfig.HealthCheckPeriod = defaultHealthCheckPeriod
	dbConfig.ConnConfig.ConnectTimeout = defaultConnectTimeout
	dbConfig.ConnConfig.Tracer = &pgxotel.QueryTracer{
		Name: "obs-test",
	}

	return dbConfig, nil
}
