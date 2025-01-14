package db

import (
	"database/sql"
	"expense_tracker/internal/db/auth"
	"expense_tracker/internal/db/expense"
	"fmt"
	_ "github.com/lib/pq"
	"log/slog"
	"os"
	"strconv"
)

type DB struct {
	logger *slog.Logger
	pg     *sql.DB
	*auth.Auth
	*expense.Expense
}

func New(logger *slog.Logger) (*DB, error) {
	pgSql, err := NewPgSQL()
	if err != nil {
		return nil, err
	}

	return &DB{
		logger:  logger,
		pg:      pgSql,
		Auth:    auth.New(logger, pgSql),
		Expense: expense.New(logger, pgSql),
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

	pgInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	db, err := sql.Open("postgres", pgInfo)
	if err != nil {
		return nil, err
	}

	return db, nil
}
