package crypto

import "app/crypto/services"

func GetService() services.CryptoService {
	return &services.BcryptCryptoService{}
}