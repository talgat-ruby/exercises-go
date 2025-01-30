package database

type CreateExerciseReq interface {
	GetUserID() string
	GetName() string
	GetDescription() string
}

type CreateExerciseResp interface{}
