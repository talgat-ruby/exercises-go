package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

// getConnection создает соединение с базой данных PostgreSQL
func getConnection() (*sql.DB, error) {
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	dbHost := "localhost"
	dbPort := 5432

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

// executeQuery выполняет SQL-запрос и возвращает результаты
func executeQuery(query string, args ...interface{}) (*sql.Rows, error) {
	db, err := getConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	return rows, nil
}

// executeQueryRow выполняет SQL-запрос и возвращает одну строку
func executeQueryRow(query string, args ...interface{}) *sql.Row {
	db, err := getConnection()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	return db.QueryRow(query, args...)
}

// initDB создает таблицу todos, если она не существует
func initDB() error {
	db, err := getConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	query := `
    CREATE TABLE IF NOT EXISTS todos (
        id SERIAL PRIMARY KEY,
        description TEXT NOT NULL,
        done BOOLEAN NOT NULL,
        created_at TIMESTAMP NOT NULL DEFAULT NOW(),
        updated_at TIMESTAMP NOT NULL DEFAULT NOW()
    );
    `

	_, err = db.Exec(query)
	return err
}
