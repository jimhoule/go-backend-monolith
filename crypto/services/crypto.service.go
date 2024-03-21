package services

type CryptoService interface {
	GenerateHashedPassword(password string) (string, error)
	ComparePassword(hashedPassword string, password string) (bool, error)
}