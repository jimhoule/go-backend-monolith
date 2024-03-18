package repositories

import "app/accounts/domain/models"

type AccountsRepository interface {
	FindAll() ([]*models.Account, error)
	FindById(id string) (*models.Account, error)
	Save(accountModel *models.Account) (*models.Account, error)
}