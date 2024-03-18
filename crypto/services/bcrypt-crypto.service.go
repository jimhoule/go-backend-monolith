package services

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

type BcryptCryptoService struct{}

func (bcs *BcryptCryptoService) GenerateHashedPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Panic(err)
	}

	return string(hashedPassword)
}

func (bcs *BcryptCryptoService) ComparePassword(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		log.Panic(err)
	}

    return err == nil
}