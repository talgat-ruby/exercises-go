package service

import (
	"errors"
	"math"
	"tracker/internal/models"
)

func (s *serviceExpence) ServiceNewUser(newUser models.NewUser) (models.NewUserResponse, error) {
	if newUser.Amount < 0 {
		newUser.Amount = int(math.Abs(float64(newUser.Amount)))
	}
	if newUser.UserID == 0 || newUser.Amount <= 0 {
		return models.NewUserResponse{}, errors.New("user not found")
	}
	return s.dbExpence.NewUserDB(newUser)
}
