package service

import (
	"errors"
	"math"
	"os"
	"tracker/internal/auth"
	"tracker/internal/models"
)

func (s *serviceExpence) ServiceNewUser(newUser models.NewUser, authToken string) (models.NewUserResponse, error) {
	if newUser.Amount < 0 {
		newUser.Amount = math.Abs(newUser.Amount)
	}

	userData, err := auth.ParseToken(authToken, os.Getenv("SECRET_TOKEN"))
	if newUser.Amount <= 0 || userData.ID == 0 {
		return models.NewUserResponse{}, errors.New("user not found")
	}
	if err != nil {
		return models.NewUserResponse{}, err
	}

	return s.dbExpence.NewUserDB(newUser, userData.ID)
}
