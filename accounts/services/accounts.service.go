package services

import (
	"app/accounts/domain/factories"
	"app/accounts/domain/models"
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

func (as *AccountsService) Create(firstName string, lastName string, email string, password string) (*models.Account, error) {
	account := as.AccountsFactory.Create(
		firstName,
		lastName,
		email,
		password,
	)

	return as.AccountsRepository.Create(account)
}