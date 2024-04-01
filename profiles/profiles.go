package profiles

import (
	"app/database/postgres"
	"app/profiles/application/services"
	"app/profiles/domain/factories"
	"app/profiles/persistence/postgres/repositories"
	"app/profiles/presenters/http/controllers"
	"app/router"
	"app/uuid"
)

func GetService(db *postgres.Db) *services.ProfilesService {
	return &services.ProfilesService{
		ProfilesFactory: &factories.ProfilesFactory{
			UuidService: uuid.GetService(),
		},
		ProfilesRepository: &repositories.PostgresProfilesRepository{
			Db: db,
		},
	}
}

func Init(mainRouter *router.MainRouter, db *postgres.Db) {
	profilesController := &controllers.ProfilesController{
		ProfilesService: GetService(db),
	}

	mainRouter.Get("/profiles/account/{accountId}", profilesController.FindAllByAccountId)
	mainRouter.Get("/profiles/{id}", profilesController.FindById)
	mainRouter.Post("/profiles", profilesController.Create)
	mainRouter.Patch("/profiles/{id}", profilesController.Update)
	mainRouter.Delete("/profiles/{id}", profilesController.Delete)
}