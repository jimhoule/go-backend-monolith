package languages

import (
	"app/database/postgres"
	"app/languages/application/services"
	"app/languages/domain/factories"
	"app/languages/persistence/postgres/repositories"
	"app/languages/presenters/http/controllers"
	"app/router"
	"app/uuid"
)

func GetService(db *postgres.Db) *services.LanguagesService {
	return &services.LanguagesService{
		LanguagesFactory: &factories.LanguagesFactory{
			UuidService: uuid.GetService(),
		},
		LanguagesRepository: &repositories.PostgresLanguagesRepository{
			Db: db,
		},
	}
}

func Init(mainRouter *router.MainRouter, db *postgres.Db) {
	languagesController := &controllers.LanguagesController{
		LanguagesService: GetService(db),
	}

	 mainRouter.Get("/languages", languagesController.FindAll)
	 mainRouter.Get("/languages/{id}", languagesController.FindById)
	 mainRouter.Post("/languages", languagesController.Create)
	 mainRouter.Put("/languages/{id}", languagesController.Update)
	 mainRouter.Delete("/languages/{id}", languagesController.Delete)
}