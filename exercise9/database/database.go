package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"alinurmyrzakhanov/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dbPath := "myworkout.db"
	if envPath := os.Getenv("DB_PATH"); envPath != "" {
		dbPath = envPath
	}

	dsn := fmt.Sprintf("%s", dbPath)
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Не удалось подключиться к БД:", err)
	}

	err = db.AutoMigrate(
		&models.User{},
		&models.Exercise{},
		&models.Workout{},
		&models.WorkoutExercise{},
	)
	if err != nil {
		log.Fatal("Ошибка миграции:", err)
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	DB = db
}
