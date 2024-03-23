package services

import (
	"app/accounts/application/payloads"
	"app/accounts/application/ports"
	"app/accounts/domain/factories"
	"app/accounts/domain/models"
)

type AccountsService struct {
	AccountsFactory factories.AccountsFactory
	AccountsRepository ports.AccountsRepository
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

func (as *AccountsService) Create(createAccountPayload payloads.CreateAccountPayload) (*models.Account, error) {
	account := as.AccountsFactory.Create(
		createAccountPayload.FirstName,
		createAccountPayload.LastName,
		createAccountPayload.Email,
		createAccountPayload.Password,
		createAccountPayload.PlanId,
	)

	return as.AccountsRepository.Create(account)
}