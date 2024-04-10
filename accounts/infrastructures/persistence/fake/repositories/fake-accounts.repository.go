package repositories

import (
	"app/accounts/domain/models"
	"fmt"
)

var accounts []*models.Account = []*models.Account{}

func ResetFakeAccountsRepository() {
	accounts = []*models.Account{}
}

type FakeAccountsRepository struct{}

func (far *FakeAccountsRepository) FindAll() ([]*models.Account, error) {
	return accounts, nil
}

func (far *FakeAccountsRepository) FindById(id string) (*models.Account, error) {
	for _, account := range accounts {
		if account.Id == id {
			return account, nil
		}
	}

	return nil, fmt.Errorf("the account with id %s does not exist", id)
}

func (far *FakeAccountsRepository) FindByEmail(email string) (*models.Account, error) {
	for _, account := range accounts {
		if account.Email == email {
			return account, nil
		}
	}

	return nil, fmt.Errorf("the account with email %s does not exist", email)
}

func (far *FakeAccountsRepository) Create(account *models.Account) (*models.Account, error) {
	accounts = append(accounts, account);

	return account, nil
}