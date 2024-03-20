package services

import (
	"app/accounts/domain/factories"
	"app/accounts/domain/models"
	"app/accounts/dtos"
	"app/accounts/persistence/repositories"
)

type AccountsService struct {
	AccountsFactory factories.AccountsFactory
	AccountsRepository repositories.AccountsRepository
}

func (as *AccountsService) FindAll() ([]*models.Account, error) {
	return as.AccountsRepository.FindAll()
}

func (as *AccountsService) FindById(id string) (*models.Account, error) {
	return as.AccountsRepository.FindById(id)
}

func (as *AccountsService) FindByEmail(email string) (*models.Account, error) {
	return as.AccountsRepository.FindByEmail(email)
}

func (as *AccountsService) Create(createAccountDto dtos.CreateAccountDto) (*models.Account, error) {
	account := as.AccountsFactory.Create(
		createAccountDto.FirstName,
		createAccountDto.LastName,
		createAccountDto.Email,
		createAccountDto.Password,
	)

	return as.AccountsRepository.Create(account)
}