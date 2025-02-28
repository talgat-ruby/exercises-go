package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

// InitDB инициализирует подключение к базе данных
func InitDB() (*sql.DB, error) {
	dbURL := os.Getenv("DB_URL") // Берём URL из переменной окружения
	if dbURL == "" {
		log.Fatal("❌ Ошибка: переменная окружения DB_URL не задана")
	}

	fmt.Println("📡 Подключение к БД...")
	database, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("❌ Ошибка подключения к БД: %v", err)
	}

	// Проверяем соединение
	err = database.Ping()
	if err != nil {
		log.Fatalf("❌ Ошибка при пинге БД: %v", err)
	}

	fmt.Println("✅ Подключение к PostgreSQL успешно!")

	// Создаём таблицу, если её нет
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS posts (
		id SERIAL PRIMARY KEY,
		title TEXT NOT NULL,
		content TEXT NOT NULL,
		category TEXT NOT NULL,
		tags TEXT[]
	);
	`
	_, err = database.Exec(createTableQuery)
	if err != nil {
		log.Fatalf("❌ Ошибка создания таблицы: %v", err)
	}

	fmt.Println("✅ Таблица 'posts' проверена или создана!")

	DB = database
	return DB, nil
}
