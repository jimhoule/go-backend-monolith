package repositories

import (
	"app/accounts/domain/models"
	"app/accounts/persistence/entities"
	"app/accounts/persistence/mappers"
	"fmt"
)

type FakeAccountsRepository struct{
	AccountsMapper mappers.AccountsMapper
}

var accountEntities []*entities.Account = []*entities.Account{}

func (far *FakeAccountsRepository) FindAll() ([]*models.Account, error) {
	var accountModels []*models.Account = []*models.Account{}
	for _, accountEntity := range accountEntities {
		acountModel := far.AccountsMapper.ToDomainModel(accountEntity)
		accountModels = append(accountModels, acountModel)
	}

	return accountModels, nil
}

func (far *FakeAccountsRepository) FindById(id string) (*models.Account, error) {
	for _, accountEntity := range accountEntities {
		if accountEntity.Id == id {
			return far.AccountsMapper.ToDomainModel(accountEntity), nil
		}
	}

	return nil, fmt.Errorf("the account with id %s does not exist", id)
}

func (far *FakeAccountsRepository) Save(accountModel *models.Account) (*models.Account, error) {
	accountEntity := far.AccountsMapper.ToEntity(accountModel)
	accountEntities = append(accountEntities, accountEntity);

	return accountModel, nil
}