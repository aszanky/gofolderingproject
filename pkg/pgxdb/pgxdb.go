package pgxdb

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewConnection(connStr string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		return nil, fmt.Errorf("gagal membuat koneksi pool baru: %w", err)
	}

	if err := pool.Ping(context.Background()); err != nil {
		pool.Close()
		return nil, fmt.Errorf("gagal melakukan ping ke database: %w", err)
	}

	return pool, nil
}
