package service

import (
	"errors"
	"log/slog"
	"math"
	"tracker/internal/models"
)

func (s *serviceExpence) ServiceEditUser(editUser models.EditUser) error {
	if editUser.IdUser <= 0 {
		slog.Error("incorrect id")
		return errors.New("incorrect id")
	}
	if editUser.Expense == 0 {
		slog.Error("consumprion not specified")
		return errors.New("consumprion not specified")
	}
	if editUser.Expense < 0 {
		editUser.Expense = math.Abs(editUser.Expense)

	}
	return s.dbExpence.DBEditUser(editUser)
}
