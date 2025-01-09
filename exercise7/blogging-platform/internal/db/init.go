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
    description TEXT NOT NULL,
    posterUrl TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
)
`

	if _, err := db.pg.Exec(stmt); err != nil {
		log.ErrorContext(ctx, "fail create table post", "error", err)
		return err
	}

	seedStmt := `
INSERT INTO post (title, description, posterUrl)
VALUES ('Lord of the Rings', 'Lord of the Rings', 'https://www.amazon.com/Lord-Rings-Movie-Poster-24x36/dp/B07D96K2QK'),
       ('Back to the future', 'Back to the future', 'https://www.amazon.com/Back-Future-Movie-Poster-Regular/dp/B001CDQF8A'),
       ('I, Robot', 'I, Robot', 'https://www.cinematerial.com/posts/i-robot-i343818');
`

	if _, err := db.pg.Exec(seedStmt); err != nil {
		log.ErrorContext(ctx, "fail seed table post", "error", err)
		return err
	}

	log.InfoContext(ctx, "success create table post")
	return nil
}
