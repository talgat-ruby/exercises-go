package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/talgat-ruby/exercises-go/exercise3/problem1/internal/env"
)

type Model struct {
	db *sql.DB
}

func NewDB(env *env.Env) *Model {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		env.DBHost, env.DBPort, env.DBUser, env.DBPass, env.DBName,
	)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	return &Model{
		db: db,
	}
}

func (m *Model) Close() {
	m.db.Close()
}
