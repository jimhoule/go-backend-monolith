package ports

import "app/accounts/domain/models"

type AccountsRepositoryPort interface {
	FindAll() ([]*models.Account, error)
	FindById(id string) (*models.Account, error)
	FindByEmail(email string) (*models.Account, error)
	Create(account *models.Account) (*models.Account, error)
}