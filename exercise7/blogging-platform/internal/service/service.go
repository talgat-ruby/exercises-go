package service

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/internal/model"
	"strings"
	"time"
)

type PostService struct {
	db *sql.DB
}

func NewPostService(db *sql.DB) *PostService {
	return &PostService{
		db: db,
	}
}

func (ps *PostService) GetPosts() ([]model.Post, error) {
	rows, err := ps.db.Query("SELECT id, title, content, category, tags, created_at, updated_at FROM posts")
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve posts: %w", err)
	}
	defer rows.Close()

	var posts []model.Post
	for rows.Next() {
		var post model.Post
		var tags string
		if err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.Category, &tags, &post.CreatedAt, &post.UpdatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan post: %w", err)
		}

		// Convert comma-separated tags string to a slice
		post.Tags = splitTags(tags)

		posts = append(posts, post)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over posts: %w", err)
	}

	return posts, nil
}

func (ps *PostService) GetPost(term string) ([]model.Post, error) {
	// Construct the query with a wildcard search if a term is provided
	query := `SELECT id, title, content, category, tags, created_at, updated_at 
			  FROM posts WHERE title ILIKE $1 OR content ILIKE $1 OR category ILIKE $1`

	// Execute the query
	rows, err := ps.db.Query(query, "%"+term+"%")
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve posts: %w", err)
	}
	defer rows.Close()

	var posts []model.Post
	for rows.Next() {
		var post model.Post
		var tags string
		if err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.Category, &tags, &post.CreatedAt, &post.UpdatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan post: %w", err)
		}

		// Convert comma-separated tags string to a slice
		post.Tags = splitTags(tags)

		posts = append(posts, post)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over posts: %w", err)
	}

	return posts, nil
}

// CreatePost inserts a new post into the database
func (ps *PostService) CreatePost(post *model.Post) (*model.Post, error) {
	query := `
		INSERT INTO posts (title, content, category, tags) 
		VALUES ($1, $2, $3, $4) 
		RETURNING id, created_at, updated_at
	`

	tags := joinTags(post.Tags)

	err := ps.db.QueryRow(query, post.Title, post.Content, post.Category, tags).
		Scan(&post.ID, &post.CreatedAt, &post.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to create post: %w", err)
	}

	return post, nil
}

// UpdatePost updates an existing post in the database
func (ps *PostService) UpdatePost(post *model.Post) (*model.Post, error) {
	updated := time.Now()
	post.UpdatedAt = updated
	query := `
		UPDATE posts 
		SET title = $1, content = $2, category = $3, tags = $4, updated_at = $5 
		WHERE id = $6 
		RETURNING created_at
	`

	tags := joinTags(post.Tags)

	err := ps.db.QueryRow(query, post.Title, post.Content, post.Category, tags, updated, post.ID).
		Scan(&post.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("post not found")
		}
		return nil, fmt.Errorf("failed to update post: %w", err)
	}

	return post, nil
}

// DeletePost deletes a post by its ID from the database
func (ps *PostService) DeletePost(id int) error {
	result, err := ps.db.Exec("DELETE FROM posts WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("failed to delete post: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to retrieve the number of deleted rows: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("post not found")
	}
	return nil
}

// Helper function to split the comma-separated tags string into a slice of strings
func splitTags(tags string) []string {
	if tags == "" {
		return []string{}
	}
	tags = tags[1 : len(tags)-1]
	return strings.Split(tags, ",") // Split tags by commas
}

// Helper function to join a slice of strings into a comma-separated tags string
func joinTags(tags []string) string {
	return "{" + strings.Join(tags, ",") + "}" // Join tags into a single comma-separated string
}
