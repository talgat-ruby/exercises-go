package repository

import (
 "database/sql"
 "log"
 "strings"
 "time"

 "github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/models"
)

type PostRepository struct {
 db *sql.DB
}

func NewPostRepository(database *sql.DB) *PostRepository {
 return &PostRepository{db: database}
}

// Create inserts a new post into the database.
func (pr *PostRepository) Create(post *models.Post) error {
 combinedTags := strings.Join(post.Tags, ",")
 query := `INSERT INTO posts (title, content, category, tags) VALUES (?, ?, ?, ?)`
 _, err := pr.db.Exec(query, post.Title, post.Content, post.Category, combinedTags)
 if err != nil {
  log.Printf("Error creating post: %v", err)
 }
 return err
}

// GetPostByID retrieves a post by its ID from the database.
func (pr *PostRepository) GetPostByID(id int) (*models.Post, error) {
 query := `SELECT id, title, content, created_at, category, tags, updated_at FROM posts WHERE id = ?`
 row := pr.db.QueryRow(query, id)

 post := &models.Post{}
 var tagsField sql.NullString
 var createdTime, updatedTime string

 err := row.Scan(&post.ID, &post.Title, &post.Content, &createdTime, &post.Category, &tagsField, &updatedTime)
 if err != nil {
  if err == sql.ErrNoRows {
   log.Printf("No post found with ID %d", id)
   return nil, nil
  }
  log.Printf("Error retrieving post by ID: %v", err)
  return nil, err
 }

 post.CreatedAt, err = time.Parse("2006-01-02 15:04:05", createdTime)
 if err != nil {
  log.Printf("Error parsing created_at: %v", err)
  return nil, err
 }

 post.UpdatedAt, err = time.Parse("2006-01-02 15:04:05", updatedTime)
 if err != nil {
  log.Printf("Error parsing updated_at: %v", err)
  return nil, err
 }

 if tagsField.Valid {
  post.Tags = strings.Split(tagsField.String, ",")
 }

 return post, nil
}

// GetAll retrieves all posts from the database.
func (pr *PostRepository) GetAll() ([]models.Post, error) {
 query := `SELECT id, title, content, created_at, category, tags, updated_at FROM posts`
 rows, err := pr.db.Query(query)
 if err != nil {
  log.Printf("Error retrieving posts: %v", err)
  return nil, err
 }
 defer rows.Close()

 var posts []models.Post
 for rows.Next() {
  var p models.Post
  var tagsField sql.NullString
  var createdTime, updatedTime string

  err := rows.Scan(&p.ID, &p.Title, &p.Content, &createdTime, &p.Category, &tagsField, &updatedTime)
  if err != nil {
   log.Printf("Error scanning post: %v", err)
   return nil, err
  }

  p.CreatedAt, err = time.Parse("2006-01-02 15:04:05", createdTime)
  if err != nil {
   log.Printf("Error parsing created_at: %v", err)
   return nil, err
  }

  p.UpdatedAt, err = time.Parse("2006-01-02 15:04:05", updatedTime)
  if err != nil {
   log.Printf("Error parsing updated_at: %v", err)
   return nil, err
  }

  if tagsField.Valid {
   p.Tags = strings.Split(tagsField.String, ",")
  }

  posts = append(posts, p)
 }

 if err := rows.Err(); err != nil {
  log.Printf("Error processing rows: %v", err)
  return nil, err
 }

 return posts, nil
}

// Update modifies an existing post in the database.
func (pr *PostRepository) Update(post *models.Post) error {
 combinedTags := strings.Join(post.Tags, ",")
 query := `UPDATE posts SET title = ?, content = ?, category = ?, tags = ?, updated_at = CURRENT_TIMESTAMP WHERE id = ?`
 _, err := pr.db.Exec(query, post.Title, post.Content, post.Category, combinedTags, post.ID)
 if err != nil {
  log.Printf("Error updating post with ID %d: %v", post.ID, err)
 }
 return err
}

// Delete removes a post from the database by its ID.
func (pr *PostRepository) Delete(id int) error {
 query := `DELETE FROM posts WHERE id = ?`
 _, err := pr.db.Exec(query, id)
 if err != nil {
  log.Printf("Error deleting post with ID %d: %v", id, err)
 }
 return err
}