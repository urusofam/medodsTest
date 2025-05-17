package storage

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	errTokenNotFound = errors.New("Token not found")
	errTokenInvalid  = errors.New("Token is invalid")
)

type Storage struct {
	Pool *pgxpool.Pool
}

func NewStorage(databaseUrl string) (*Storage, error) {
	pool, err := pgxpool.New(context.Background(), databaseUrl)
	if err != nil {
		return nil, fmt.Errorf("error creating storage: %w", err)
	}

	if err = pool.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &Storage{Pool: pool}, nil
}
