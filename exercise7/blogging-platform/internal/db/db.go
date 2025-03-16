package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

// Создаем собственный тип DB, который будет оберткой для *sql.DB
type DB struct {
	*sql.DB
}

// Функция для создания нового подключения к базе данных с типом DB
func New() (*DB, error) {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	sqlDB, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return &DB{sqlDB}, nil
}

// Функция для удаления поста по ID
func (d *DB) DeletePostByID(id int) error {
	// Логика удаления записи из базы данных по id
	_, err := d.Exec("DELETE FROM posts WHERE id = $1", id)
	return err
}
