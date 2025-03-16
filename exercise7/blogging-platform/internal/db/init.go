package db

import (
	"context"
)

func (db *DB) Init(ctx context.Context) error {
	log := db.logger.With("method", "Init")
	stmt := `
CREATE TABLE IF NOT EXISTS post (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		content TEXT NOT NULL,
		author_id INTEGER NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
`
	if _, err := db.pg.Exec(stmt); err != nil {
		log.ErrorContext(ctx, "fail create table post", "error", err)
		return err
	}
	log.InfoContext(ctx, "success create table post")
	return nil
}
