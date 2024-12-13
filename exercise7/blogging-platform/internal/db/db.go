package db

import (
	"database/sql"
	"fmt"
	conf "github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/internal/config"
)

type ConfDB struct {
	conf *conf.Config
	db   *sql.DB
}

func New(conf *conf.Config) (*ConfDB, error) {
	db, err := NewDb(
		conf.DB.Host,
		conf.DB.Port,
		conf.DB.User,
		conf.DB.Password,
		conf.DB.DBName,
	)
	if err != nil {
		return nil, err
	}

	return &ConfDB{
		conf: conf,
		db:   db,
	}, nil
}

func NewDb(host string, port int, user string, password string, dbname string) (*sql.DB, error) {
	fmt.Println(host, port, user, password, dbname)
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
