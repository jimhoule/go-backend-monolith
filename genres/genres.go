package genres

import (
	"app/database/postgres"
	"app/genres/application/services"
	"app/genres/domain/factories"
	"app/genres/persistence/postgres/repositories"
	"app/genres/presenters/http/controllers"
	"app/router"
	"app/translations"
	"app/uuid"
)

func GetService(db *postgres.Db) *services.GenresService {
	return &services.GenresService{
		GenresFactory: &factories.GenresFactory{
			UuidService: uuid.GetService(),
		},
		GenresRepository: &repositories.PostgresGenresRepository{
			Db: db,
		},
		TranslationsService: translations.GetService(db),
	}
}

func Init(mainRouter *router.MainRouter, db *postgres.Db) {
	genresController := &controllers.GenresController{
		GenresService: GetService(db),
	}

	mainRouter.Get("/genres", genresController.FindAll)
	mainRouter.Get("/genres/{id}", genresController.FindById)
	mainRouter.Post("/genres", genresController.Create)
}