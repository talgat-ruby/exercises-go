package service

import (
	"tracker/internal/db"
	"tracker/internal/models"
)

type ServiceExpence interface {
	ServiceNewUser(newUser models.NewUser) (models.NewUserResponse, error)
}

type serviceExpence struct {
	dbExpence db.ExpencesDBSt
}

func NewServiceExpence(dataBase db.ExpencesDBSt) ServiceExpence {
	return &serviceExpence{dbExpence: dataBase}
}
