package db

import (
	"database/sql"
	conf "github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/internal/config"
)

type DB struct {
	conf *conf.Config
	db   *sql.DB
}

func New(conf *conf.Config) (*DB, error) {
	return &DB{}, nil
}
