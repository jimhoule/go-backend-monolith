package repositories

import "app/accounts/domain/models"

type AccountsRepository interface {
	FindAll() ([]*models.Account, error)
	FindById(id string) (*models.Account, error)
	FindByEmail(email string) (*models.Account, error)
	Create(account *models.Account) (*models.Account, error)
}