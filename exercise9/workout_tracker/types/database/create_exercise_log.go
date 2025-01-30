package database

import "time"

type CreateExerciseLogReq interface {
	GetWorkoutDate() time.Time
	GetWorkoutTime() time.Time
	SetNumber() int
	Repetitions() int
	Weight() float64
	WeightUnit() string
	Distance() float64
	DistanceType() string
	Notes() string
}

type CreateExerciseLogResp{}
