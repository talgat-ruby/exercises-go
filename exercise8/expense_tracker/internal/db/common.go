package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"
)

type ExpencesDB interface{
	
}

type expencesDB struct {
	newDb *sql.DB
}

func NewExpenceDB(db *sql.DB) ExpencesDB {
	return &expencesDB{
		newDb: db,
	}
}

func Init() (error, *sql.DB) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_NAME")

	info := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, database)
	log.Println("Waiting for database to be ready...")
	for i := 0; i < 5; i++ {
		newDb, err := sql.Open("postgres", info)
		if err == nil {
			if err = newDb.Ping(); err == nil {
				log.Println("Succesfully connected to the database!")
			}
			return nil, newDb

		}
		log.Printf("Database not ready yet ... retrying(%d/5)\n", i+1)
		time.Sleep(5 * time.Second)
	}
	return fmt.Errorf("failed to connect to database"), nil
}
