package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

type BlogPost struct {
	ID        int       `json:"id"`
	Title     string    `json:"title" binding:"required"`
	Content   string    `json:"content" binding:"required"`
	Category  string    `json:"category" binding:"required"`
	Tags      []string  `json:"tags" binding:"required"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

var db *sql.DB

func main() {
	var err error
	connStr := "postgres://astanamac:adlet@localhost:5432/blog?sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Ошибка при подключении к PostgreSQL:", err)
	}

	// Проверка соединения
	err = db.Ping()
	if err != nil {
		log.Fatal("Не удалось подключиться к базе данных:", err)
	} else {
		fmt.Println("✅ Успешное подключение к базе данных!")
	}

	r := gin.Default()
	r.POST("/posts", createPost)
	r.PUT("/posts/:id", updatePost)
	r.DELETE("/posts/:id", deletePost)
	r.GET("/posts/:id", getPost)
	r.GET("/posts", getAllPosts)

	fmt.Println("🚀 Сервер запущен на порту 8080")
	r.Run(":8080")
}

func createPost(c *gin.Context) {
	var post BlogPost
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	timeNow := time.Now()
	post.CreatedAt, post.UpdatedAt = timeNow, timeNow

	query := `INSERT INTO posts (title, content, category, tags, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	err := db.QueryRow(query, post.Title, post.Content, post.Category, pq.Array(post.Tags), post.CreatedAt, post.UpdatedAt).Scan(&post.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, post)
}

func updatePost(c *gin.Context) {
	id := c.Param("id")
	var post BlogPost
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post.UpdatedAt = time.Now()
	query := `UPDATE posts SET title=$1, content=$2, category=$3, tags=$4, updated_at=$5 WHERE id=$6`
	result, err := db.Exec(query, post.Title, post.Content, post.Category, pq.Array(post.Tags), post.UpdatedAt, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	c.JSON(http.StatusOK, post)
}

func deletePost(c *gin.Context) {
	id := c.Param("id")
	query := `DELETE FROM posts WHERE id=$1`
	result, err := db.Exec(query, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	c.Status(http.StatusNoContent)
}

func getPost(c *gin.Context) {
	id := c.Param("id")
	var post BlogPost
	var tags pq.StringArray

	query := `SELECT id, title, content, category, tags, created_at, updated_at FROM posts WHERE id=$1`
	err := db.QueryRow(query, id).Scan(&post.ID, &post.Title, &post.Content, &post.Category, (*pq.StringArray)(&tags), &post.CreatedAt, &post.UpdatedAt)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	post.Tags = tags
	c.JSON(http.StatusOK, post)
}

func getAllPosts(c *gin.Context) {
	searchTerm := c.Query("term")
	query := `SELECT id, title, content, category, tags, created_at, updated_at FROM posts WHERE title ILIKE $1 OR content ILIKE $1 OR category ILIKE $1`
	rows, err := db.Query(query, "%"+searchTerm+"%")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var posts []BlogPost
	for rows.Next() {
		var post BlogPost
		var tags pq.StringArray

		if err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.Category, (*pq.StringArray)(&tags), &post.CreatedAt, &post.UpdatedAt); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		post.Tags = tags
		posts = append(posts, post)
	}

	c.JSON(http.StatusOK, posts)
}
