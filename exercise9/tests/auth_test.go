package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"alinurmyrzakhanov/database"
	"alinurmyrzakhanov/models"
	"alinurmyrzakhanov/repositories"
	"alinurmyrzakhanov/routes"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Не удалось открыть in-memory DB: %v", err)
	}

	err = db.AutoMigrate(
		&models.User{},
		&models.Exercise{},
		&models.Workout{},
		&models.WorkoutExercise{},
	)
	if err != nil {
		t.Fatalf("Ошибка миграции: %v", err)
	}

	database.DB = db
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	routes.SetupRoutes(r)
	return r
}

func createTestExercise(t *testing.T, name, category string) *models.Exercise {
	ex := models.Exercise{
		Name:        name,
		Description: "test desc",
		Category:    category,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	err := repositories.CreateExercise(&ex)
	if err != nil {
		t.Fatalf("Не удалось создать упражнение в тесте: %v", err)
	}
	return &ex
}

func TestCreateExercise(t *testing.T) {
	setupTestDB(t)
	router := setupRouter()

	body := `{
		"name": "Bench Press",
		"description": "Chest exercise",
		"category": "strength"
	}`
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/exercises", bytes.NewBuffer([]byte(body)))
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code, "Должен вернуться статус 201")

	var created models.Exercise
	err := json.Unmarshal(w.Body.Bytes(), &created)
	assert.Nil(t, err)

	assert.Equal(t, "Bench Press", created.Name)
	assert.Equal(t, "strength", created.Category)
	assert.NotZero(t, created.ID, "ID должен быть сгенерирован")
}

func TestGetAllExercises(t *testing.T) {
	setupTestDB(t)
	router := setupRouter()

	createTestExercise(t, "Push-up", "strength")
	createTestExercise(t, "Squat", "strength")

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/exercises", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Статус должен быть 200")

	var exercises []models.Exercise
	err := json.Unmarshal(w.Body.Bytes(), &exercises)
	assert.Nil(t, err)

	assert.GreaterOrEqual(t, len(exercises), 2, "Должно вернуться как минимум 2 упражнения")
}

func TestGetExerciseByID(t *testing.T) {
	setupTestDB(t)
	router := setupRouter()

	testEx := createTestExercise(t, "Push-up", "strength")

	w := httptest.NewRecorder()

	url := "/exercises/" + strconv.Itoa(int(testEx.ID))
	req, _ := http.NewRequest("GET", url, nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var fetched models.Exercise
	err := json.Unmarshal(w.Body.Bytes(), &fetched)
	assert.Nil(t, err)

	assert.Equal(t, testEx.ID, fetched.ID)
	assert.Equal(t, "Push-up", fetched.Name)
}

func TestUpdateExercise(t *testing.T) {
	setupTestDB(t)
	router := setupRouter()

	testEx := createTestExercise(t, "Bench Press", "strength")

	updateBody := `{
		"name": "Bench Press (updated)",
		"description": "New desc",
		"category": "upper body"
	}`
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT",
		"/exercises/"+(func() string {
			return string(rune(testEx.ID + '0'))
		})(),
		bytes.NewBuffer([]byte(updateBody)),
	)
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var updated models.Exercise
	err := json.Unmarshal(w.Body.Bytes(), &updated)
	assert.Nil(t, err)

	assert.Equal(t, testEx.ID, updated.ID)
	assert.Equal(t, "Bench Press (updated)", updated.Name)
	assert.Equal(t, "New desc", updated.Description)
	assert.Equal(t, "upper body", updated.Category)
}

func TestDeleteExercise(t *testing.T) {
	setupTestDB(t)
	router := setupRouter()

	testEx := createTestExercise(t, "Plank", "core")

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE",
		"/exercises/"+(func() string {
			return string(rune(testEx.ID + '0'))
		})(),
		nil,
	)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code, "Ожидаем 204 No Content")

	w2 := httptest.NewRecorder()
	req2, _ := http.NewRequest("GET",
		"/exercises/"+(func() string {
			return string(rune(testEx.ID + '0'))
		})(),
		nil,
	)
	router.ServeHTTP(w2, req2)

	assert.Equal(t, http.StatusNotFound, w2.Code, "После удаления GET должен вернуть 404")
}

func TestGetNonExistentExercise(t *testing.T) {
	setupTestDB(t)
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/exercises/9999", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code, "Ожидаем 404 для несуществующего ID")
}

func TestCreateExercise_ValidationError(t *testing.T) {
	setupTestDB(t)
	router := setupRouter()
	body := `{
		"description": "Missing name",
		"category": "strength"
	}`
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/exercises", bytes.NewBuffer([]byte(body)))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code, "Ожидаем 400 Bad Request")
}
