package ports

import "app/profiles/domain/models"

type ProfilesRepositoryPort interface {
	FindAllByAccountId(accountId string) ([]*models.Profile, error)
	FindById(id string) (*models.Profile, error)
	Update(id string, profile *models.Profile) (*models.Profile, error)
	Delete(id string) (string, error)
	Create(profile *models.Profile) (*models.Profile, error)
}