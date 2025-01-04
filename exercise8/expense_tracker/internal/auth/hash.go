package auth

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
)

const (
	SaltLength   = 16
	PepperLength = 32
	HashLength   = 64
	Iterations   = 210000
)

func HashPassword(pasword, pepper string) (string, string, error) {
	salt := make([]byte, SaltLength)
	_, err := rand.Read(salt)
	if err != nil {
		return "", "", fmt.Errorf("error generation salt: %w", err)
	}
	hash := hashWithSaltAndPepper([]byte(pasword), salt, []byte(pepper))
	hashStr := base64.StdEncoding.EncodeToString(hash)
	saltStr := base64.StdEncoding.EncodeToString(salt)
	return hashStr, saltStr, nil
}

func hashWithSaltAndPepper(password, salt, pepper []byte) []byte {
	pepperedPassword := append(password, pepper...)
	hash := hmac.New(sha512.New, salt)
	result := pepperedPassword
	for i := 0; i < Iterations; i++ {
		hash.Reset()
		hash.Write(result)
		result = hash.Sum(nil)
	}
	return result
}

// func VerifyPassword(password, salt, pepper, hash string)( bool, error) {

// }
