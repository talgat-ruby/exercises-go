package todo

import (
	"database/sql"
	"fmt"
	"log/slog"

	_ "github.com/lib/pq"
	"github.com/talgat-ruby/exercises-go/exercise3/problem1/config"
	todoRepository "github.com/talgat-ruby/exercises-go/exercise3/problem1/internal/repositories/todo"
	"github.com/talgat-ruby/exercises-go/exercise3/problem1/internal/repositories/todo/postgresql"
)

func NewRepository(conf *config.Config) todoRepository.Repository {
	slog.Info("hello from new Repository")
	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		conf.Database.Host, conf.Database.Port, conf.Database.User,
		conf.Database.Password, conf.Database.DbName, conf.Database.SSLMode,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return postgresql.NewRepository(db)
}
