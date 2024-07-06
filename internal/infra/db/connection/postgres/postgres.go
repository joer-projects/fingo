package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func NewConnection(ctx context.Context) (*pgx.Conn, error) {
	db, err := pgx.Connect(ctx, "user=admin password=admin host=localhost port=5432 dbname=accounting sslmode=disable")
	// TODO: Config env vars for db connection

	if err != nil {
		return db, err
	}

	return db, nil
}
