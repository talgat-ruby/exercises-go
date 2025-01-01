package api

import (
	"context"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/internal/db"
	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/internal/models"
)

type API struct {
	app *fiber.App
	db  *db.DB
}

func New() *API {
	return &API{
		app: fiber.New(),
	}
}

func (a *API) Start(ctx context.Context) error {
	a.routes()
	return a.app.Listen(":8080")
}

func (a *API) routes() {
	a.app.Get("/", a.handleRoot)
	a.app.Post("/posts", a.createPost)
	a.app.Put("/posts/:id", a.updatePost)
	a.app.Delete("/posts/:id", a.deletePost)
	a.app.Get("/posts/:id", a.getPost)
	a.app.Get("/posts", a.getAllPosts)
}
func (a *API) handleRoot(c *fiber.Ctx) error {
	return c.SendString("Welcome to the Blogging Platform!")
}

// CRUD Handlers
func (a *API) createPost(c *fiber.Ctx) error {
	var post struct {
		Title    string   `json:"title"`
		Content  string   `json:"content"`
		Category string   `json:"category"`
		Tags     []string `json:"tags"`
	}

	if err := c.BodyParser(&post); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid data")
	}

	// Вставка нового поста в базу данных
	query := `INSERT INTO posts (title, content, category, tags) 
              VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at`
	var newPost struct {
		ID        int       `json:"id"`
		CreatedAt time.Time `json:"createdAt"`
		UpdatedAt time.Time `json:"updatedAt"`
	}

	err := a.db.QueryRow(query, post.Title, post.Content, post.Category, post.Tags).Scan(&newPost.ID, &newPost.CreatedAt, &newPost.UpdatedAt)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to create post")
	}

	return c.Status(fiber.StatusCreated).JSON(newPost)
}

func (a *API) updatePost(c *fiber.Ctx) error {

	id := c.Params("id")
	var post struct {
		Title    string   `json:"title"`
		Content  string   `json:"content"`
		Category string   `json:"category"`
		Tags     []string `json:"tags"`
	}

	if err := c.BodyParser(&post); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid data")
	}

	// Обновление поста в базе данных
	query := `UPDATE posts 
              SET title = $1, content = $2, category = $3, tags = $4, updated_at = CURRENT_TIMESTAMP
              WHERE id = $5`
	_, err := a.db.Exec(query, post.Title, post.Content, post.Category, post.Tags, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to update post")
	}

	return c.Status(fiber.StatusOK).SendString("Post updated successfully")
}

func (a *API) deletePost(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	if err := a.db.DeletePostByID(id); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete post"})
	}

	return c.SendStatus(http.StatusNoContent)
}

func (a *API) getPost(c *fiber.Ctx) error {

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	post := models.Post{
		ID:        id,
		Title:     "Sample Post",
		Content:   "Sample Content",
		Category:  "Technology",
		Tags:      []string{"Sample", "Tech"},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return c.Status(http.StatusOK).JSON(post)
}

func (a *API) getAllPosts(c *fiber.Ctx) error {
	term := c.Query("term")

	posts := []models.Post{
		{
			ID:        1,
			Title:     "Sample Post",
			Content:   "Sample Content",
			Category:  "Technology",
			Tags:      []string{"Sample", "Tech"},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	if term != "" {

		term = strings.ToLower(term)
		filtered := []models.Post{}
		for _, post := range posts {
			if strings.Contains(strings.ToLower(post.Title), term) ||
				strings.Contains(strings.ToLower(post.Content), term) ||
				strings.Contains(strings.ToLower(post.Category), term) {
				filtered = append(filtered, post)
			}
		}
		return c.Status(http.StatusOK).JSON(filtered)
	}

	return c.Status(http.StatusOK).JSON(posts)
}
