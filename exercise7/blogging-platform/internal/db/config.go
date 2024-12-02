package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"log/slog"

	_ "github.com/lib/pq"
)

func InitDatabase(cfg map[string]string, ctx context.Context) *sql.DB {
	// Setup the connection to the database
	dbConn, err := setupConnection(cfg)
	if err != nil {
		slog.ErrorContext(ctx, "db initialize error", "service", "db", "error", err)
		panic(err) // Panic on connection error
	}

	// Perform database migration
	errMigration := migrate(dbConn)
	if errMigration != nil {
		slog.ErrorContext(ctx, "db migration error", "service", "db", "error", err)
		panic(errMigration) // Panic on migration error
	}

	// Return the initialized connection
	return dbConn
}

func migrate(db *sql.DB) error {
	// Simple migration script to create the posts table if it doesn't exist
	migrationScript := `
	CREATE TABLE IF NOT EXISTS posts (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    category VARCHAR(100) NOT NULL,
    tags TEXT[],  -- Using an array of text for tags
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
`
	_, err := db.Exec(migrationScript)
	if err != nil {
		return fmt.Errorf("failed to initialize the database: %w", err)
	}

	log.Println("Database migration performed successfully.")
	return nil
}

func setupConnection(cfg map[string]string) (*sql.DB, error) {
	// Construct the DSN (Data Source Name)
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg["DB_HOST"], cfg["DB_PORT"], cfg["DB_USER"],
		cfg["DB_PASSWORD"], cfg["DB_NAME"], cfg["DB_SSLMODE"],
	)

	// Open the database connection
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Check if the connection is valid
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("Database initialized successfully.")
	return db, nil
}
