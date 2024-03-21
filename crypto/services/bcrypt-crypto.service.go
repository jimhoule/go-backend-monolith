package services

import (
	"golang.org/x/crypto/bcrypt"
)

type BcryptCryptoService struct{}

func (bcs *BcryptCryptoService) GenerateHashedPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func (bcs *BcryptCryptoService) ComparePassword(hashedPassword string, password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false, err
	}

    return true, nil
}