package db

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"os"
	"strconv"

	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/internal/db/category"
	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/internal/db/post"
	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/internal/db/tag"

	_ "github.com/lib/pq"
)

type DB struct {
	logger *slog.Logger
	pg     *sql.DB
	*category.Category
	*post.Post
	*tag.Tag
}

func New(logger *slog.Logger) (*DB, error) {
	pgsql, err := newPgSQL()
	if err != nil {
		return nil, err
	}

	return &DB{
		logger:   logger,
		pg:       pgsql,
		Category: category.New(pgsql, logger),
		Post:     post.New(pgsql, logger),
		Tag:      tag.New(pgsql, logger),
	}, nil
}

func newPgSQL() (*sql.DB, error) {
	host := os.Getenv("DB_HOST")
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		return nil, err
	}
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (db *DB) Ping(ctx context.Context) error {
	err := db.pg.PingContext(ctx)
	if err != nil {
		db.logger.ErrorContext(ctx, "failed to connect to database", "error", err)
		return err
	}

	db.logger.InfoContext(ctx, "success connected to database")
	return nil
}
