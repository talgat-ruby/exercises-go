package db

import (
	"database/sql"
	"fmt"
	"log/slog"
	"os"
	"strconv"

	_ "github.com/lib/pq"
	"github.com/webaiz/exercise7/blogging-platform/internal/db/auth"
	"github.com/webaiz/exercise7/blogging-platform/internal/db/post"
)

type DB struct {
	logger *slog.Logger
	pg     *sql.DB
	*post.Post
	*auth.Auth
}

func New(logger *slog.Logger) (*DB, error) {
	pgsql, err := NewPgSQL()
	if err != nil {
		return nil, err
	}

	return &DB{
		logger: logger,
		pg:     pgsql,
		Post:   post.New(pgsql, logger),
		Auth:   auth.New(pgsql, logger),
	}, nil
}

func NewPgSQL() (*sql.DB, error) {
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
