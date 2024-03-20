package accounts

import (
	"app/accounts/controllers"
	"app/accounts/domain/factories"
	"app/accounts/persistence/mappers"
	"app/accounts/persistence/repositories"
	"app/accounts/services"
	"app/database/postgres"
	"app/router"

	cryptoService "app/crypto/services"
	uuidService "app/uuid/services"
)

func GetService(db *postgres.Db) *services.AccountsService {
	return &services.AccountsService{
		AccountsFactory: factories.AccountsFactory{
			UuidService: &uuidService.NativeUuidService{},
			CryptoService: &cryptoService.BcryptCryptoService{},
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