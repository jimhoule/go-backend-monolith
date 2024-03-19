package mappers

import (
	"app/accounts/domain/models"
	"app/accounts/persistence/entities"
)

type AccountsMapper struct {}

func (am *AccountsMapper) ToDomainModel(accountEntity *entities.Account) *models.Account {
	return &models.Account{
		Id: accountEntity.Id,
		FirstName: accountEntity.FirstName,
		LastName: accountEntity.LastName,
		Email: accountEntity.Email,
		Password: accountEntity.Password,
	}
}

func (am *AccountsMapper) ToEntity(accountModel *models.Account) *entities.Account {
	return &entities.Account{
		Id: accountModel.Id,
		FirstName: accountModel.FirstName,
		LastName: accountModel.LastName,
		Email: accountModel.Email,
		Password: accountModel.Password,
	}
}