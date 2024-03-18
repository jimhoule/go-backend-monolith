package services

import (
	"app/accounts/domain/factories"
	"app/accounts/domain/models"
	"app/accounts/dtos"
	"app/accounts/persistence/repositories"
)

type AccountsService struct {
	AccountsRepository repositories.AccountsRepository
}

func (as *AccountsService) FindAll() ([]*models.Account, error) {
	return as.AccountsRepository.FindAll()
}

func (as *AccountsService) FindById(id string) (*models.Account, error) {
	return as.AccountsRepository.FindById(id)
}

func (as *AccountsService) Save(createAccountDto dtos.CreateAccountDto) (*models.Account, error) {
	account := factories.CreateAccount(
		createAccountDto.FirstName,
		createAccountDto.LastName,
		createAccountDto.Email,
		createAccountDto.Password,
	)

	return as.AccountsRepository.Save(account)
}