package factories

import (
	"app/accounts/domain/models"
	cryptoService "app/crypto/services"
	uuidService "app/uuid/services"
)

var nativeUuidservice = uuidService.NativeUuidService{}
var brcyptCryptoService = cryptoService.BcryptCryptoService{}

func CreateAccount(firstName string, lastName string, email string, password string) *models.Account {
	return &models.Account{
		Id: nativeUuidservice.Generate(),
		FirstName: firstName,
		LastName: lastName,
		Email: email,
		Password: brcyptCryptoService.GenerateHashedPassword(password),
	}
}