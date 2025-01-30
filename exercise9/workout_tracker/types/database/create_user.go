package database

import "time"

type CreateUserReq interface {
	GetUsername() string
	GetEmail() string
	GetPasswordHash() string
	GetRole() string
	GetSalt() string
}

type CreateUserResp interface {
	GetID() string
	GetUsername() string
	GetEmail() string
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
}
