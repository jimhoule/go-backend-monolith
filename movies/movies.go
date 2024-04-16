package movies

import (
	"app/aws"
	"app/database/postgres"
	moviesServices "app/movies/application/services"
	"app/movies/domain/factories"
	"app/movies/infrastructures/persistence/postgres/repositories"
	"app/movies/infrastructures/storage"
	"app/movies/presenters/http/controllers"
	"app/router"
	"app/transactions"
	"app/translations"
	"app/uuid"
	"os"
)

func GetService(db *postgres.Db) *moviesServices.MoviesService {
	return &moviesServices.MoviesService{
		MoviesFactory: &factories.MoviesFactory{
			UuidService: uuid.GetService(),
		},
		MoviesRepository: &repositories.PostgresMoviesRepository{
			Db: db,
		},
		MoviesStorage: &storage.S3Storge{
			S3Service: aws.CreateS3Service(os.Getenv("AWS_VIDEO_UPLOADS_BUCKET_NAME")),
		},
		TranslationsService: translations.GetService(db),
		TransactionsService: transactions.GetService(db),
	}
}

func Init(mainRouter *router.MainRouter, db *postgres.Db) {
	moviesController := &controllers.MoviesController{
		MoviesService: GetService(db),
	}

	mainRouter.Get("/movies", moviesController.FindAll)
	mainRouter.Get("/movies/{id}", moviesController.FindById)
	mainRouter.Post("/movies", moviesController.Create)
	mainRouter.Post("/movies/uploads", moviesController.Upload)
	mainRouter.Put("/movies/{id}", moviesController.Update)
	mainRouter.Delete("/movies/{id}", moviesController.Delete)
}