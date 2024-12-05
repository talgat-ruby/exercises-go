package db

import (
	"context"
)

func (db *DB) Init(ctx context.Context) error {
	log := db.logger.With("method", "Init")

	if err := db.InitCategory(ctx); err != nil {
		return err
	}

	if err := db.InitTag(ctx); err != nil {
		return err
	}

	if err := db.InitPost(ctx); err != nil {
		return err
	}

	if err := db.InitPostTags(ctx); err != nil {
		return err
	}

	log.InfoContext(ctx, "success create tables for blogging system")
	return nil
}

func (db *DB) InitCategory(ctx context.Context) error {
	log := db.logger.With("table", "category")
	stmt := `
	CREATE TABLE IF NOT EXISTS category 
	(
		id serial PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
		updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
	);`

	if _, err := db.pg.Exec(stmt); err != nil {
		log.ErrorContext(ctx, "failed to create category table", "error", err)
		return err
	}

	return nil
}

func (db *DB) InitTag(ctx context.Context) error {
	log := db.logger.With("table", "tag")
	stmt := `
	CREATE TABLE IF NOT EXISTS tag
	(
		id serial PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
		updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
	);`

	if _, err := db.pg.Exec(stmt); err != nil {
		log.ErrorContext(ctx, "failed to create tag table", "error", err)
		return err
	}

	return nil
}

func (db *DB) InitPost(ctx context.Context) error {
	log := db.logger.With("table", "post")
	stmt := `
	CREATE TABLE IF NOT EXISTS post
	(
		id serial PRIMARY KEY,
		id_category int references category (id) on delete cascade not null,
		title VARCHAR(255) NOT NULL,
		content TEXT NOT NULL,
		created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
		updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
	);`

	if _, err := db.pg.Exec(stmt); err != nil {
		log.ErrorContext(ctx, "failed to create post table", "error", err)
		return err
	}

	return nil
}

func (db *DB) InitPostTags(ctx context.Context) error {
	log := db.logger.With("table", "post_tags")

	stmt := `
	CREATE TABLE IF NOT EXISTS post_tags
	(
		id serial PRIMARY KEY,
		id_post int references post (id) on delete cascade not null,
		id_tag int references tag (id) on delete cascade not null,
		created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
		updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
	);`

	if _, err := db.pg.Exec(stmt); err != nil {
		log.ErrorContext(ctx, "failed to create post_tags table", "error", err)
		return err
	}

	return nil
}
