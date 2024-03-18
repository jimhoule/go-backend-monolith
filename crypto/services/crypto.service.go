package services

type CryptoService interface {
	GenerateHashedPassword(password string) string
	ComparePassword(hashedPassword string, password string) bool
}