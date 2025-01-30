package database

type CreateWorkoutPlanReq interface {
	GetUserID() string
	GetName() string
	GetDescription() string
	GetActivityStatus() bool
}

type CreateWorkoutPlanResp interface {
	GetName() string
	GetActivityStatus() bool
}
