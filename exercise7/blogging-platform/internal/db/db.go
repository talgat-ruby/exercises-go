package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/models"
)

func New() (*sql.DB, error) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	exists, err := CheckTable(db)
	if err != nil {
		return nil, err
	}
	if !exists {
		res, err := ReadSQLFIle("file_sql/create_table")
		if err != nil {
			return nil, err
		}
		_, err = db.Exec(res)
		if err != nil {
			return nil, err
		}
	}

	fmt.Println("Database connected successfully")
	return db, nil

}

func NewPostCreator(body models.Blog) error {
	db, err := New()
	if err != nil {
		return err
	}
	res, err := ReadSQLFIle("file_sql/create_new_post")
	if err != nil {
		return err
	}
	_, err = db.Exec(res, body.Title, body.Content, body.Category, body.Tags, body.CreatedAt, body.UpdatedAt)
	return err
}
func ReadSQLFIle(path string) (string, error) {
	res, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(res), nil
}
func CheckTable(db *sql.DB) (bool, error) {
	query, err := ReadSQLFIle("file_sql/check_table")
	if err != nil {
		return false, err
	}
	var exists bool
	err = db.QueryRow(query, "posts").Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}
func DBGetPosts() ([]models.Blog, error) {
	db, err := New()
	if err != nil {
		return nil, err
	}
	var res []models.Blog
	// res, err := db.Exec("SELECT  FROM posts")
	rows, err := db.Query("SELECT * FROM posts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var post models.Blog
		err := rows.Scan(&post.Id, &post.Title, &post.Content, &post.Category, &post.Tags, &post.CreatedAt, &post.UpdatedAt)
		if err != nil {
			return nil, err
		}
		res = append(res, post)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return res, nil
}
