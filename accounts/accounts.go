package accounts

import (
	"app/accounts/application/services"
	"app/accounts/domain/factories"
	"app/accounts/infrastructures/persistence/postgres/repositories"
	"app/accounts/presenters/http/controllers"
	"app/crypto"
	"app/database"
	"app/router"
	"app/uuid"
)

func GetService(db *database.Db) *services.AccountsService {
	return &services.AccountsService{
		AccountsFactory: &factories.AccountsFactory{
			UuidService: uuid.GetService(),
			CryptoService: crypto.GetService(),
		},
		AccountsRepository: &repositories.PostgresAccountsRepository{
			Db: db,
		},
	}
}

func Init(mainRouter *router.MainRouter, db *database.Db) {
	accountsController := &controllers.AccountsController{
		AccountsService: GetService(db),
	}

	mainRouter.Get("/accounts", accountsController.FindAll)
	mainRouter.Get("/accounts/{id}", accountsController.FindById)
}