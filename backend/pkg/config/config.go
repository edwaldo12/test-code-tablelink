package config

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

func NewPostgresPool() (*pgxpool.Pool, error) {
	dsn := os.Getenv("postgresql://postgres:edwaldo@localhost:5432/tablelink_test?schema=public")
	if dsn == "" {
		dsn = "postgres://postgres:edwaldo@localhost:5432/tablelink_test"
	}
	cfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to parse pgxpool config: %w", err)
	}

	pool, err := pgxpool.ConnectConfig(context.TODO(), cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	return pool, nil
}
