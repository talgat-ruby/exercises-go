package database

import (
	"log"
	"time"

	"alinurmyrzakhanov/models"
)

func SeedExercises() {
	var count int64
	DB.Model(&models.Exercise{}).Count(&count)
	if count > 0 {
		log.Println("Exercises уже существуют, сидер пропущен.")
		return
	}

	exercises := []models.Exercise{
		{Name: "Push-up", Description: "Classic push-up for chest", Category: "strength", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Squat", Description: "Bodyweight squat for legs", Category: "strength", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Running", Description: "Running cardio exercise", Category: "cardio", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Plank", Description: "Core strength exercise", Category: "core", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}

	if err := DB.Create(&exercises).Error; err != nil {
		log.Println("Ошибка при создании упражнений:", err)
	} else {
		log.Println("Exercises успешно засеяны в БД.")
	}
}
