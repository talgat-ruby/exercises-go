package controllers

import (
	"net/http"
	"strconv"
	"time"

	"alinurmyrzakhanov/models"
	"alinurmyrzakhanov/repositories"

	"github.com/gin-gonic/gin"
)

func CreateExercise(c *gin.Context) {
	var input struct {
		Name        string `json:"name" binding:"required"`
		Description string `json:"description"`
		Category    string `json:"category" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ex := models.Exercise{
		Name:        input.Name,
		Description: input.Description,
		Category:    input.Category,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := repositories.CreateExercise(&ex); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось создать упражнение"})
		return
	}

	c.JSON(http.StatusCreated, ex)
}

func GetExercises(c *gin.Context) {
	exercises, err := repositories.GetAllExercises()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении упражнений"})
		return
	}
	c.JSON(http.StatusOK, exercises)
}

func GetExerciseByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный ID"})
		return
	}

	ex, err := repositories.GetExerciseByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Упражнение не найдено"})
		return
	}

	c.JSON(http.StatusOK, ex)
}

func UpdateExercise(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный ID"})
		return
	}

	existing, err := repositories.GetExerciseByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Упражнение не найдено"})
		return
	}

	var input struct {
		Name        string `json:"name" binding:"required"`
		Description string `json:"description"`
		Category    string `json:"category" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	existing.Name = input.Name
	existing.Description = input.Description
	existing.Category = input.Category
	existing.UpdatedAt = time.Now()

	if err := repositories.UpdateExercise(existing); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обновлении упражнения"})
		return
	}

	c.JSON(http.StatusOK, existing)
}

func DeleteExercise(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный ID"})
		return
	}

	if err := repositories.DeleteExercise(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при удалении упражнения"})
		return
	}

	c.Status(http.StatusNoContent)
}
