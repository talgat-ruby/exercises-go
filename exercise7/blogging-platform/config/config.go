package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() *sql.DB {
	dsn := "user:root@tcp(127.0.0.1:3306)/blog_platform"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("DB is not responding: %v", err)
	}

	fmt.Println("Successful connection to the database")
	return db
}
