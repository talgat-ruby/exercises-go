package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	conf "github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/internal/config"
	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/internal/db/blog"
	"log/slog"
	"os"
	"strconv"
)

type ConfDB struct {
	Conf *conf.Config
	DB   *sql.DB
	*blog.Blogs
}

func New(conf *conf.Config, logger *slog.Logger) (*ConfDB, error) {
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
		Conf:  conf,
		DB:    db,
		Blogs: blog.NewBlog(logger, db),
	}, nil
}

func NewDb(host string, port int, user string, password string, dbname string) (*sql.DB, error) {
	fmt.Println(host, port, user, password, dbname)
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		return nil, err
	}
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
