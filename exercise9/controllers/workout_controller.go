package controllers

import (
	"net/http"
	"strconv"
	"time"

	"alinurmyrzakhanov/models"
	"alinurmyrzakhanov/repositories"

	"github.com/gin-gonic/gin"
)

func CreateWorkout(c *gin.Context) {
	userID := GetUserID(c)
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Нет user_id"})
		return
	}

	var input struct {
		Title            string `json:"title" binding:"required"`
		Scheduled        string `json:"scheduled"`
		Comment          string `json:"comment"`
		WorkoutExercises []struct {
			ExerciseID uint   `json:"exerciseId" binding:"required"`
			Sets       int    `json:"sets" binding:"required"`
			Reps       int    `json:"reps" binding:"required"`
			Weight     int    `json:"weight"`
			Comment    string `json:"comment"`
		} `json:"exercises"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var scheduled time.Time
	if input.Scheduled != "" {
		t, err := time.Parse(time.RFC3339, input.Scheduled)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный формат времени. Используйте RFC3339."})
			return
		}
		scheduled = t
	}

	workout := models.Workout{
		UserID:    userID,
		Title:     input.Title,
		Scheduled: scheduled,
		Comment:   input.Comment,
		IsDone:    false,
	}
	for _, wex := range input.WorkoutExercises {
		workout.WorkoutExercises = append(workout.WorkoutExercises, models.WorkoutExercise{
			ExerciseID: wex.ExerciseID,
			Sets:       wex.Sets,
			Reps:       wex.Reps,
			Weight:     wex.Weight,
			Comment:    wex.Comment,
		})
	}

	if err := repositories.CreateWorkout(&workout); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании тренировки"})
		return
	}
	c.JSON(http.StatusCreated, workout)
}

func UpdateWorkout(c *gin.Context) {
	userID := GetUserID(c)
	idParam := c.Param("id")
	workoutID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный ID"})
		return
	}

	existing, err := repositories.GetWorkoutByID(uint(workoutID), userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Тренировка не найдена"})
		return
	}

	var input struct {
		Title            string `json:"title" binding:"required"`
		Scheduled        string `json:"scheduled"`
		Comment          string `json:"comment"`
		IsDone           bool   `json:"isDone"`
		WorkoutExercises []struct {
			ExerciseID uint   `json:"exerciseId" binding:"required"`
			Sets       int    `json:"sets" binding:"required"`
			Reps       int    `json:"reps" binding:"required"`
			Weight     int    `json:"weight"`
			Comment    string `json:"comment"`
		} `json:"exercises"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	existing.Title = input.Title
	existing.Comment = input.Comment
	existing.IsDone = input.IsDone

	if input.Scheduled != "" {
		t, err := time.Parse(time.RFC3339, input.Scheduled)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный формат времени"})
			return
		}
		existing.Scheduled = t
	}
	var newWex []models.WorkoutExercise
	for _, wex := range input.WorkoutExercises {
		newWex = append(newWex, models.WorkoutExercise{
			ExerciseID: wex.ExerciseID,
			Sets:       wex.Sets,
			Reps:       wex.Reps,
			Weight:     wex.Weight,
			Comment:    wex.Comment,
		})
	}
	existing.WorkoutExercises = newWex

	if err := repositories.UpdateWorkout(existing); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обновлении"})
		return
	}

	c.JSON(http.StatusOK, existing)
}

func DeleteWorkout(c *gin.Context) {
	userID := GetUserID(c)
	idParam := c.Param("id")
	workoutID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный ID"})
		return
	}

	if err := repositories.DeleteWorkout(uint(workoutID), userID); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Тренировка не найдена или ошибка при удалении"})
		return
	}
	c.Status(http.StatusNoContent)
}

func ListWorkouts(c *gin.Context) {
	userID := GetUserID(c)
	onlyPending := c.Query("pending") == "true"

	workouts, err := repositories.ListWorkouts(userID, onlyPending)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении списка"})
		return
	}
	c.JSON(http.StatusOK, workouts)
}

func GetWorkoutReport(c *gin.Context) {
	userID := GetUserID(c)
	fromStr := c.Query("from")
	toStr := c.Query("to")

	if fromStr == "" || toStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Параметры ?from=YYYY-MM-DD&to=YYYY-MM-DD обязательны"})
		return
	}

	from, errFrom := time.Parse("2006-01-02", fromStr)
	to, errTo := time.Parse("2006-01-02", toStr)
	if errFrom != nil || errTo != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат дат (используйте YYYY-MM-DD)"})
		return
	}
	if from.After(to) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Дата 'from' не может быть позже 'to'"})
		return
	}

	workouts, err := repositories.GetWorkoutsInRange(userID, from, to)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при формировании отчёта"})
		return
	}

	totalWorkouts := len(workouts)

	c.JSON(http.StatusOK, gin.H{
		"from":         fromStr,
		"to":           toStr,
		"totalRecords": totalWorkouts,
		"workouts":     workouts,
	})
}
