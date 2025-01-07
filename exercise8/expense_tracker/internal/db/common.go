package db

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
)

type ExpencesDB interface {
	Init() (error, *sql.DB)
}

type ExpencesDBSt struct {
	NewDb *sql.DB
}

func NewExpenceDB() (*ExpencesDBSt, error) {
	err, db := Newpostgresql()
	if err != nil {
		return &ExpencesDBSt{}, err
	}
	if db == nil {
		return nil, errors.New("bd is nil")
	}
	return &ExpencesDBSt{
		NewDb: db,
	}, nil
}

func Newpostgresql() (error, *sql.DB) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_NAME")

	info := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, database)
	log.Println("Waiting for database to be ready...")
	fmt.Println("info", info)
	newDb, err := sql.Open("postgres", info)
	if err != nil {
		return err, nil
	}
	return nil, newDb
}
