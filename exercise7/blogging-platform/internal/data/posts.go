package data

import (
	"database/sql"
	"errors"
	"github.com/lib/pq"
	"time"
)

type Post struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Category  string    `json:"category"`
	Tags      []string  `json:"tags"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PostModel struct {
	DB *sql.DB
}

func (p PostModel) GetAll() ([]*Post, error) {
	posts := make([]*Post, 0)

	query := `
		SELECT id, title, content, category, tags, created_at, updated_at
		FROM posts`

	rows, err := p.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var post Post

		if err := rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.Category,
			pq.Array(&post.Tags),
			&post.CreatedAt,
			&post.UpdatedAt,
		); err != nil {
			return nil, err
		}

		posts = append(posts, &post)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

func (p PostModel) Insert(post *Post) error {
	query := `
		INSERT INTO posts (title, content, category, tags)
		VALUES ($1, $2, $3, $4)
		RETURNING id, title, content, category, tags, created_at, updated_at`

	args := []interface{}{post.Title, post.Content, post.Category, pq.Array(post.Tags)}

	row := p.DB.QueryRow(query, args...)
	if err := row.Err(); err != nil {
		return err
	}
	if err := row.Scan(
		&post.ID,
		&post.Title,
		&post.Content,
		&post.Category,
		pq.Array(&post.Tags),
		&post.CreatedAt,
		&post.UpdatedAt,
	); err != nil {
		return err
	}
	return nil
}

func (p PostModel) Get(id int64) (*Post, error) {

	query := `
		SELECT id, title, content, category, tags, created_at, updated_at
		FROM posts
		WHERE id = $1`

	row := p.DB.QueryRow(query, id)
	if err := row.Err(); err != nil {

		return nil, err
	}
	var post Post

	if err := row.Scan(
		&post.ID,
		&post.Title,
		&post.Content,
		&post.Category,
		pq.Array(&post.Tags),
		&post.CreatedAt,
		&post.UpdatedAt,
	); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &post, nil
}

func (p PostModel) Update(post *Post) error {
	query := `
		UPDATE posts
		SET title = $1, content = $2, category = $3, tags = $4, updated_at = now()
		WHERE id = $5
		RETURNING updated_at`

	args := []interface{}{
		post.Title,
		post.Content,
		post.Category,
		pq.Array(post.Tags),
		post.ID,
	}

	err := p.DB.QueryRow(query, args...).Scan(&post.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (p PostModel) Delete(id int64) error {
	query := `DELETE FROM posts WHERE id = $1`

	result, err := p.DB.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrRecordNotFound
	}

	return nil
}
