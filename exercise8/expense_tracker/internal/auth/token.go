package auth

import (
	"errors"
	"fmt"
	"strconv"
	"time"

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
		fmt.Println("token is nil")
		return data, fmt.Errorf("invalid token")
	case t.Valid:
		claims, ok := t.Claims.(jwt.MapClaims)
		if !ok {
			fmt.Println("token is nil1")
			return data, fmt.Errorf("invalid token")
		}
		id, err := claims.GetSubject()
		if err != nil {
			fmt.Println(token, secret)
			fmt.Println("get subject")
			return data, fmt.Errorf("invalid token")
		}
		intID, err := strconv.Atoi(id)
		if err != nil {
			fmt.Println("atoi")
			return data, fmt.Errorf("invalid token")
		}
		email, ok := claims["email"].(string)
		if !ok {
			fmt.Println("claims email")
			return data, fmt.Errorf("invalid token")
		}
		data = &models.UserData{
			ID:    intID,
			Email: email,
		}
		return data, nil
	case errors.Is(err, jwt.ErrTokenMalformed):
		fmt.Println("malformed")
		return data, fmt.Errorf("invalid token")
	case errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet):
		fmt.Println("expired")
		return data, err

	case errors.Is(err, jwt.ErrTokenSignatureInvalid):
		fmt.Println("signature")
		return data, fmt.Errorf("invalid signature")
	default:
		return data, nil
	}
}

func GenerateToken(user *models.UserData, secret string) (*models.Tokens, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = user.Email
	claims["sub"] = user.ID
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return nil, fmt.Errorf("error generating token: %w", err)
	}
	refeshToken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refeshToken.Claims.(jwt.MapClaims)
	rtClaims["sub"] = user.ID
	rtClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	rt, err := refeshToken.SignedString([]byte(secret))
	if err != nil {
		return nil, fmt.Errorf("error generating token: %w", err)
	}
	return &models.Tokens{
		AccessToken:  t,
		RefreshToken: rt,
	}, nil
}
