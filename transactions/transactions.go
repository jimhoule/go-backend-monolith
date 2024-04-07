package transactions

import (
	"app/database/postgres"
	"app/transactions/application/services"
	"app/transactions/persistence/postgres/repositories"
)

func GetService(db *postgres.Db) *services.TransactionsService {
	return &services.TransactionsService{
		TransactionsRepository: &repositories.PostgresTransactionsRepository{
			Db: db,
		},
	}
}