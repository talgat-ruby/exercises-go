package service

import (
	"errors"
	"tracker/internal/models"
)

func (s *serviceExpence) ServiceGetExpenses(id int) (models.TransactionsReponse, error) {
	if id <= 0 {
		return models.TransactionsReponse{}, errors.New("incorrect id")
	}

	return s.dbExpence.DBGetExpenses(id)
}
