package storage

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage struct {
	pool *pgxpool.Pool
}

func NewStorage(databaseUrl string) (*Storage, error) {
	pool, err := pgxpool.New(context.Background(), databaseUrl)
	if err != nil {
		return nil, fmt.Errorf("error creating storage: %w", err)
	}

	if err = pool.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &Storage{pool: pool}, nil
}

func (s *Storage) Close() {
	s.pool.Close()
}

func (s *Storage) CreateUser(ctx context.Context, GUID string) error {
	query := "INSERT INTO users (guid) VALUES ($1)"
	_, err := s.pool.Exec(ctx, query, GUID)
	if err != nil {
		return fmt.Errorf("error creating user: %w", err)
	}
	return nil
}

func (s *Storage) SaveRefreshToken(
	ctx context.Context,
	userGUID string,
	tokenHash string,
	userAgent string,
	IP string,
) error {
	query := `INSERT INTO refresh_tokens
			(user_guid, token_hash, user_agent, ip)
			VALUES ($1, $2, $3, $4)`
	_, err := s.pool.Exec(ctx, query, userGUID, tokenHash, userAgent, IP)
	if err != nil {
		return fmt.Errorf("error creating refresh token: %w", err)
	}
	return nil
}

func (s *Storage) GetRefreshToken(ctx context.Context, userGUID string) (string, error) {
	query := `SELECT user_agent
			  FROM refresh_tokens
			  WHERE user_guid = $1`
	tokenHash := ""
	row := s.pool.QueryRow(ctx, query, userGUID)
	err := row.Scan(&tokenHash)
	if err != nil {
		return "", fmt.Errorf("error getting refresh token: %w", err)
	}
	return tokenHash, nil
}

func (s *Storage) DeleteRefreshToken(ctx context.Context, tokenHash string) error {
	query := `DELETE FROM refresh_tokens WHERE token_hash = $1`
	_, err := s.pool.Exec(ctx, query, tokenHash)
	if err != nil {
		return fmt.Errorf("error deleting refresh token: %w", err)
	}
	return nil
}

func (s *Storage) RefreshToken(
	ctx context.Context,
	oldUserGUID string,
	oldTokenHash string,
	newTokenHash string,
	newUserAgent string,
	newIP string,
) error {
	currUserAgent := ""
	query := `SELECT user_agent
			  FROM refresh_tokens
			  WHERE user_guid = $1 AND token_hash = $2`
	row := s.pool.QueryRow(ctx, query, oldUserGUID, oldTokenHash)
	err := row.Scan(&currUserAgent)
	if err != nil {
		return fmt.Errorf("token not found: %w", err)
	}

	err = s.DeleteRefreshToken(ctx, oldTokenHash)
	if err != nil {
		return fmt.Errorf("error deleting old refresh token: %w", err)
	}

	query = `INSERT INTO refresh_tokens 
            (user_id, token_hash, user_agent, ip)
        	VALUES ($1, $2, $3, $4)`
	_, err = s.pool.Exec(ctx, query, oldUserGUID, newTokenHash, newUserAgent, newIP)
	if err != nil {
		return fmt.Errorf("error creating new refresh token: %w", err)
	}
	return nil
}
