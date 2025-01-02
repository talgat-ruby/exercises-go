package auth

import (
	"errors"
	"fmt"

	"tracker/internal/models"

	"github.com/golang-jwt/jwt/v5"
)

func ParseToken(token, secret string) (*models.UserData, error) {
	t, err := jwt.Parse(
		token, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		},
	)
	var data *models.UserData
	switch {
	case t == nil:
		return data, fmt.Errorf("invalid token")
	case t.Valid:
		claims, ok := t.Claims.(jwt.MapClaims)
		if !ok {
			return data, fmt.Errorf("invalid token")
		}
		id, err := claims.GetSubject()
		if err != nil {
			return data, fmt.Errorf("invalid token")
		}
		email, ok := claims["email"].(string)
		if !ok {
			return data, fmt.Errorf("invalid token")
		}
		data = &models.UserData{
			ID:    id,
			Email: email,
		}
		return data, nil
	case errors.Is(err, jwt.ErrTokenMalformed):
		return data, fmt.Errorf("invalid token")
	case errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet):
		return data, err
	case errors.Is(err, jwt.ErrTokenSignatureInvalid):
		return data, fmt.Errorf("invalid signature")
	default:
		return data, nil
	}
}
