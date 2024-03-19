package repositories

import (
	"app/accounts/domain/models"
	"app/accounts/persistence/entities"
	"app/accounts/persistence/mappers"
	"app/database/postgres"
	"log"
)

type PostgresAccountsRepository struct{
	AccountsMapper mappers.AccountsMapper
	Db *postgres.Db
}

func (par *PostgresAccountsRepository) FindAll() ([]*models.Account, error) {
	var accountEntities []*entities.Account
	result := par.Db.Find(&accountEntities)
	if result.Error != nil {
		log.Panicf("Error finding all Acounts: %s", result.Error.Error())
		return nil, result.Error
	}

	var accountModels []*models.Account = []*models.Account{}
	for _, accountEntity := range accountEntities {
		acountModel := par.AccountsMapper.ToDomainModel(accountEntity)
		accountModels = append(accountModels, acountModel)
	}

	return accountModels, nil
}

func (par *PostgresAccountsRepository) FindById(id string) (*models.Account, error) {
	var accountEntities []*entities.Account
	result := par.Db.Where("Id = ?", id).Find(&accountEntities)
	if result.Error != nil {
		log.Panicf("Error finding Account with id %s: %s", id, result.Error.Error())
		return nil, result.Error
	}

	return par.AccountsMapper.ToDomainModel(accountEntities[0]), nil
}

func  (par *PostgresAccountsRepository) Create(accountModel *models.Account) (*models.Account, error) {
	accountEntity := par.AccountsMapper.ToEntity(accountModel)
	result := par.Db.Create(accountEntity)
	if result.Error != nil {
		log.Panicf("Error saving an Acount: %s", result.Error.Error())
		return nil, result.Error
	}

	return accountModel, nil
}