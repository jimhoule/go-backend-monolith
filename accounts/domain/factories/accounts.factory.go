package factories

import (
	"app/accounts/domain/models"
	cryptoService "app/crypto/services"
	uuidService "app/uuid/services"
)

type AccountsFactory struct{
	UuidService uuidService.UuidService
	CryptoService cryptoService.CryptoService
}

func (af *AccountsFactory) Create(firstName string, lastName string, email string, password string) *models.Account {
	hashedPassword, _ := af.CryptoService.GenerateHashedPassword(password)

	return &models.Account{
		Id: af.UuidService.Generate(),
		FirstName: firstName,
		LastName: lastName,
		Email: email,
		Password: hashedPassword,
	}
}