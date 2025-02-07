// controllers/post_controller.go
package controllers

import (
	"net/http"
	"strconv"
	"time"

	"alinurmyrzakhanov/models"
	"alinurmyrzakhanov/repositories"

	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {
	var input struct {
		Title    string   `json:"title" binding:"required"`
		Content  string   `json:"content" binding:"required"`
		Category string   `json:"category" binding:"required"`
		Tags     []string `json:"tags"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	post := models.Post{
		Title:     input.Title,
		Content:   input.Content,
		Category:  input.Category,
		Tags:      input.Tags,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := repositories.CreatePost(&post); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось создать пост"})
		return
	}
	c.JSON(http.StatusCreated, post)
}

func GetPost(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный ID"})
		return
	}

	post, err := repositories.GetPostByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Пост не найден"})
		return
	}

	c.JSON(http.StatusOK, post)
}

func GetAllPosts(c *gin.Context) {
	term := c.Query("term")
	if term != "" {
		posts, err := repositories.SearchPostsByTerm(term)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при поиске постов"})
			return
		}
		c.JSON(http.StatusOK, posts)
	} else {
		posts, err := repositories.GetAllPosts()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении постов"})
			return
		}
		c.JSON(http.StatusOK, posts)
	}
}

func UpdatePost(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный ID"})
		return
	}
	existingPost, err := repositories.GetPostByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Пост не найден"})
		return
	}

	var input struct {
		Title    string   `json:"title" binding:"required"`
		Content  string   `json:"content" binding:"required"`
		Category string   `json:"category" binding:"required"`
		Tags     []string `json:"tags"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	existingPost.Title = input.Title
	existingPost.Content = input.Content
	existingPost.Category = input.Category
	existingPost.Tags = input.Tags
	existingPost.UpdatedAt = time.Now()

	if err := repositories.UpdatePost(existingPost); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось обновить пост"})
		return
	}

	c.JSON(http.StatusOK, existingPost)
}

func DeletePost(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный ID"})
		return
	}

	_, getErr := repositories.GetPostByID(uint(id))
	if getErr != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Пост не найден"})
		return
	}

	if err := repositories.DeletePost(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось удалить пост"})
		return
	}

	c.Status(http.StatusNoContent)
}
