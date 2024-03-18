package repositories

import "app/accounts/domain/models"

type PostgresAccountsRepository struct{}

func (par *PostgresAccountsRepository) FindAll() ([]*models.Account, error) {
	return []*models.Account{}, nil
}

func (par *PostgresAccountsRepository) FindById(id string) (*models.Account, error) {
	account := &models.Account{
		Id: "fakeId",
		FirstName: "John",
		LastName: "Smith",
		Email: "john@example.com",
		Password: "1234",
		IsMembershipCancelled: false,
	}

	return account, nil
}

func  (par *PostgresAccountsRepository) Save(account *models.Account) (*models.Account, error) {
	return account, nil
}