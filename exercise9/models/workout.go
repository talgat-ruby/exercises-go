package models

import (
	"time"
)

type Workout struct {
	ID               uint              `gorm:"primaryKey" json:"id"`
	UserID           uint              `gorm:"index" json:"userId"`
	Title            string            `gorm:"not null" json:"title"`
	Scheduled        time.Time         `json:"scheduled"`
	Comment          string            `json:"comment"`
	IsDone           bool              `json:"isDone"`
	CreatedAt        time.Time         `json:"createdAt"`
	UpdatedAt        time.Time         `json:"updatedAt"`
	WorkoutExercises []WorkoutExercise `json:"exercises"`
}

type WorkoutExercise struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	WorkoutID  uint   `gorm:"index" json:"workoutId"`
	ExerciseID uint   `gorm:"index" json:"exerciseId"`
	Sets       int    `json:"sets"`
	Reps       int    `json:"reps"`
	Weight     int    `json:"weight"`
	Comment    string `json:"comment"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	Exercise Exercise `gorm:"foreignKey:ID;references:ExerciseID" json:"exercise"`
}
