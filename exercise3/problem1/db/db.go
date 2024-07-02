package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/talgat-ruby/exercises-go/exercise3/problem1/config"
	"log/slog"
	"os"
	"path/filepath"
)

const scriptsPath = "./resources/db"

type Model struct {
	conf *config.DBConfig
	db   *sql.DB
}

func NewModel(conf *config.DBConfig) (*Model, error) {
	db, err := NewDB(conf)

	if err != nil {
		return nil, err
	}

	m := &Model{
		conf: conf,
		db:   db,
	}

	return m, nil
}

func NewDB(conf *config.DBConfig) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		conf.SharedConfig.Host, conf.SharedConfig.Port, conf.User, conf.Password, conf.DBName,
	)

	slog.Info("Connecting to database")
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		slog.Error("Failed to connect to database")
		panic(err)
	}

	if err = db.Ping(); err != nil {
		slog.Error("Failed to ping database")
		panic(err)
	}

	scripts, err := os.ReadDir(scriptsPath)
	if err == nil {
		for _, script := range scripts {
			scriptPath := filepath.Join(scriptsPath, script.Name())
			scriptName := script.Name()
			command, err := os.ReadFile(scriptPath)
			if err != nil {
				slog.Warn("Failed to read script - " + scriptName)
				continue
			}
			_, err = db.Query(string(command))
			if err != nil {
				slog.Error("Failed to execute script - " + scriptName)
			} else {
				slog.Info("Successfully executed - " + scriptName)
			}
		}
	}

	slog.Info("Successfully connected to database")
	return db, nil
}
