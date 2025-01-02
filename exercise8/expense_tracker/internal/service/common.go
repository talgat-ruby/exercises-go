package service

import (
	"tracker/internal/db"
)

type ServiceExpence interface{}

type serviceExpence struct {
	dbExpence db.ExpencesDB
}

func NewServiceExpence(dataBase db.ExpencesDB) ServiceExpence {
	return &serviceExpence{dbExpence: dataBase}
}

