package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

const dbDSN = "host=127.0.0.1 port=54320 dbname=orders_local user=orders_user password=orders_password sslmode=disable"

type Database struct {
	pool *pgxpool.Pool
}

func New(ctx context.Context) (*Database, error) {
	pool, err := pgxpool.New(ctx, dbDSN)
	if err != nil {
		return nil, err
	}

	return &Database{
		pool: pool,
	}, nil
}

func (db *Database) Close() {
	db.pool.Close()
}
