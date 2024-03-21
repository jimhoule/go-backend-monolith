package accounts

import (
	"app/accounts/controllers"
	"app/accounts/domain/factories"
	"app/accounts/persistence/mappers"
	"app/accounts/persistence/repositories"
	"app/accounts/services"
	"app/crypto"
	"app/database/postgres"
	"app/router"
	"app/uuid"
)

func GetService(db *postgres.Db) *services.AccountsService {
	return &services.AccountsService{
		AccountsFactory: factories.AccountsFactory{
			UuidService: uuid.GetService(),
			CryptoService: crypto.GetService(),
		},
		AccountsRepository: &repositories.PostgresAccountsRepository{
			AccountsMapper: mappers.AccountsMapper{},
			Db: db,
		},
	}
}

func Init(mainRouter *router.MainRouter, db *postgres.Db) {
	accountsController := controllers.AccountsController{
		AccountsService: *GetService(db),
	}

	mainRouter.Get("/accounts", accountsController.FindAll)
	mainRouter.Get("/accounts/{id}", accountsController.FindById)

	mainRouter.Post("/accounts", accountsController.Create)
}