package db

import (
	"context"
)

func (db *DB) Init(ctx context.Context) error {
	log := db.logger.With("method", "Init")

	stmt := `
CREATE TABLE IF NOT EXISTS post (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    content TEXT NOT NULL,
    category TEXT NOT NULL,
	tags TEXT[] NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
)
`

	if _, err := db.pg.Exec(stmt); err != nil {
		log.ErrorContext(ctx, "fail create table post", "error", err)
		return err
	}

	log.InfoContext(ctx, "success create table post")
	return nil
}
