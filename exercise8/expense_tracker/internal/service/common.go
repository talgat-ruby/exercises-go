package service

import (
	"tracker/internal/db"
	"tracker/internal/models"
)

type ServiceExpence interface {
	ServiceNewUser(newUser models.NewUser, authToken string) (models.NewUserResponse, error)
	ServiceEditUser(editUser models.EditUserRequest, id string) error
	ServiceBalance(id int) (map[string]float64, error)
	ServiceGetExpenses(id int) (models.TransactionsReponse, error)
}

type serviceExpence struct {
	dbExpence db.ExpencesDBSt
}

func NewServiceExpence(dataBase db.ExpencesDBSt) ServiceExpence {
	return &serviceExpence{dbExpence: dataBase}
}
