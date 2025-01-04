package service

import (
	"tracker/internal/db"
)

type ServiceExpence interface{}

type serviceExpence struct {
	dbExpence db.ExpencesDBSt
}

func NewServiceExpence(dataBase db.ExpencesDBSt) ServiceExpence {
	return &serviceExpence{dbExpence: dataBase}
}
