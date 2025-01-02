package auth

import (
	"crypto/rand"
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
	// hash :=
}

func hashWithSaltAndPepper(password, salt, pepper []byte) []byte {
	pepperedPassword := make([]byte, len(password)+len(pepper))
	
}
