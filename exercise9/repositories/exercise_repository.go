package repositories

import (
	"alinurmyrzakhanov/database"
	"alinurmyrzakhanov/models"
)

func CreateExercise(ex *models.Exercise) error {
	return database.DB.Create(ex).Error
}

func GetExerciseByID(id uint) (*models.Exercise, error) {
	var ex models.Exercise
	err := database.DB.First(&ex, id).Error
	if err != nil {
		return nil, err
	}
	return &ex, nil
}

func GetAllExercises() ([]models.Exercise, error) {
	var exercises []models.Exercise
	err := database.DB.Find(&exercises).Error
	if err != nil {
		return nil, err
	}
	return exercises, nil
}

func UpdateExercise(ex *models.Exercise) error {
	return database.DB.Save(ex).Error
}

func DeleteExercise(id uint) error {
	return database.DB.Delete(&models.Exercise{}, id).Error
}
