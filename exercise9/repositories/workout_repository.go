package repositories

import (
	"alinurmyrzakhanov/database"
	"alinurmyrzakhanov/models"
	"time"

	"gorm.io/gorm"
)

func CreateWorkout(workout *models.Workout) error {
	return database.DB.Transaction(func(tx *gorm.DB) error {
		// Сначала создаём сам Workout
		if err := tx.Create(workout).Error; err != nil {
			return err
		}
		for i := range workout.WorkoutExercises {
			workout.WorkoutExercises[i].WorkoutID = workout.ID
			if err := tx.Create(&workout.WorkoutExercises[i]).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func GetWorkoutByID(id uint, userID uint) (*models.Workout, error) {
	var workout models.Workout
	err := database.DB.Preload("WorkoutExercises.Exercise").
		Where("id = ? AND user_id = ?", id, userID).
		First(&workout).Error
	if err != nil {
		return nil, err
	}
	return &workout, nil
}

func UpdateWorkout(workout *models.Workout) error {
	return database.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(workout).Error; err != nil {
			return err
		}
		if err := tx.Where("workout_id = ?", workout.ID).Delete(&models.WorkoutExercise{}).Error; err != nil {
			return err
		}
		for i := range workout.WorkoutExercises {
			workout.WorkoutExercises[i].WorkoutID = workout.ID
			if err := tx.Create(&workout.WorkoutExercises[i]).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func DeleteWorkout(id uint, userID uint) error {
	return database.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("id = ? AND user_id = ?", id, userID).
			Delete(&models.Workout{}).Error; err != nil {
			return err
		}
		if err := tx.Where("workout_id = ?", id).Delete(&models.WorkoutExercise{}).Error; err != nil {
			return err
		}
		return nil
	})
}

func ListWorkouts(userID uint, onlyPending bool) ([]models.Workout, error) {
	var workouts []models.Workout
	query := database.DB.Preload("WorkoutExercises.Exercise").
		Where("user_id = ?", userID)

	if onlyPending {
		now := time.Now()
		query = query.Where("scheduled >= ? OR is_done = false", now)
	}
	err := query.Order("scheduled ASC").Find(&workouts).Error
	if err != nil {
		return nil, err
	}
	return workouts, nil
}

func GetWorkoutsInRange(userID uint, from, to time.Time) ([]models.Workout, error) {
	var workouts []models.Workout
	err := database.DB.Preload("WorkoutExercises.Exercise").
		Where("user_id = ? AND scheduled BETWEEN ? AND ?", userID, from, to).
		Order("scheduled ASC").
		Find(&workouts).Error
	if err != nil {
		return nil, err
	}
	return workouts, nil
}
