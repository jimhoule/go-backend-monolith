package transactions

import (
	"app/database"
	"app/transactions/application/services"
	"app/transactions/infrastructures/persistence/postgres/repositories"
)

func GetService(db *database.Db) *services.TransactionsService {
	return &services.TransactionsService{
		TransactionsRepository: &repositories.PostgresTransactionsRepository{
			Db: db,
		},
	}
}