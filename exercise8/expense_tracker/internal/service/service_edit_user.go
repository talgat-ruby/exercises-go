package service

import (
	"errors"
	"log/slog"
	"math"
	"strconv"
	"tracker/internal/models"
)

func (s *serviceExpence) ServiceEditUser(editUser models.EditUserRequest, id string) error {
	intId, err := strconv.Atoi(id)
	if err != nil {
		slog.Error("failed to convert id to int", "error", err)
	}
	if intId <= 0 {
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
	return s.dbExpence.DBEditUser(editUser, intId)
}
