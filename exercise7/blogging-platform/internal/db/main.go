package db

import (
	"database/sql"
	"fmt"
	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/internal/db/blog"
	"log/slog"
	"os"
	"strconv"
)

type DB struct {
	logger *slog.Logger
	pg     *sql.DB
	*blog.Blog
}

func New(logger *slog.Logger) (*DB, error) {
	pgsql, err := NewPgSQL()
	if err != nil {
		return nil, fmt.Errorf("creating postgres connection: %w", err)
	}

	//Verifying database connections
	if err := pgsql.Ping(); err != nil {
		return nil, fmt.Errorf("pinging database: %w", err)
	}

	return &DB{
		logger: logger,
		pg:     pgsql,
		Blog:   blog.New(pgsql, logger),
	}, nil
}

func NewPgSQL() (*sql.DB, error) {

	config := struct {
		host     string
		port     int
		user     string
		password string
		dbname   string
	}{}

	var err error
	config.host = getEnvOrDefault("DB_HOST", "localhost")
	portStr := getEnvOrDefault("DB_PORT", "5432")
	config.port, err = strconv.Atoi(portStr)
	if err != nil {
		return nil, fmt.Errorf("invalid DB_PORT value: %w", err)
	}

	config.user = getEnvOrDefault("DB_USER", "postgres")
	config.password = os.Getenv("DB_PASSWORD")
	config.dbname = getEnvOrDefault("DB_NAME", "blogdb")

	if config.password == "" {
		return nil, fmt.Errorf("DB_PASSWORD is required")
	}

	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.host, config.port, config.user, config.password, config.dbname,
	)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, fmt.Errorf("opening database connection: %w", err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)

	return db, nil
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
