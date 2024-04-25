package genres

import (
	"app/database"
	"app/genres/application/services"
	"app/genres/domain/factories"
	"app/genres/infrastructures/persistence/postgres/repositories"
	"app/genres/presenters/http/controllers"
	"app/router"
	"app/transactions"
	"app/translations"
	"app/uuid"
)

func GetService(db *database.Db) *services.GenresService {
	return &services.GenresService{
		GenresFactory: &factories.GenresFactory{
			UuidService: uuid.GetService(),
		},
		GenresRepository: &repositories.PostgresGenresRepository{
			Db: db,
		},
		TranslationsService: translations.GetService(db),
		TransactionsService: transactions.GetService(db),
	}
}

func Init(mainRouter *router.MainRouter, db *database.Db) {
	genresController := &controllers.GenresController{
		GenresService: GetService(db),
	}

	mainRouter.Get("/genres", genresController.FindAll)
	mainRouter.Get("/genres/{id}", genresController.FindById)
	mainRouter.Post("/genres", genresController.Create)
	mainRouter.Put("/genres/{id}", genresController.Update)
	mainRouter.Delete("/genres/{id}", genresController.Delete)
}