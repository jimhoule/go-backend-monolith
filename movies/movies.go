package movies

import (
	"app/aws"
	"app/database"
	"app/gql"
	"app/movies/application/services"
	"app/movies/domain/factories"
	"app/movies/infrastructures/persistence/postgres/repositories"
	"app/movies/infrastructures/storage"
	"app/movies/presenters/gql/resolvers"
	"app/movies/presenters/http/controllers"
	"app/movies/presenters/websocket/events"
	"app/movies/presenters/websocket/handlers"
	"app/router"
	"app/transactions"
	"app/transcoder"
	"app/translations"
	"app/uuid"
	"app/websocket"
	"os"
)

func GetService(db *database.Db) *services.MoviesService {
	return &services.MoviesService{
		MoviesFactory: &factories.MoviesFactory{
			UuidService: uuid.GetService(),
		},
		MoviesRepository: &repositories.PostgresMoviesRepository{
			Db: db,
		},
		MoviesStorage: &storage.S3Storge{
			S3Service: aws.CreateS3Service(os.Getenv("AWS_VIDEO_UPLOADS_BUCKET_NAME")),
		},
		TranscoderService: transcoder.GetService(),
		TransactionsService: transactions.GetService(db),
		TranslationsService: translations.GetService(db),
	}
}

func Init(mainRouter *router.MainRouter, websocketServer *websocket.Server, gqlServer *gql.Server, db *database.Db) {
	moviesController := &controllers.MoviesController{
		MoviesService: GetService(db),
	}

	mainRouter.Get("/movies", moviesController.FindAll)
	mainRouter.Get("/movies/{id}", moviesController.FindById)
	mainRouter.Post("/movies", moviesController.Create)
	mainRouter.Post("/movies/uploads", moviesController.Upload)
	mainRouter.Put("/movies/{id}", moviesController.Update)
	mainRouter.Delete("/movies/{id}", moviesController.Delete)

	moviesHandler := &handlers.MoviesHandler{
		MoviesService: GetService(db),
	}

	websocketServer.On(
		events.StartTranscodeDashAndUploadVideo,
		moviesHandler.HandleTranscodeDashAndUploadVideo,
	)

	moviesResolver := &resolvers.MoviesResolver{
		MoviesService: GetService(db),
	}

	findAllMoviesQuery := moviesResolver.GetFindAllMoviesQuery()
	findMovieByIdQuery := moviesResolver.GetFindMovieByIdQuery()

	gqlServer.Add(findAllMoviesQuery.Name, findAllMoviesQuery)
	gqlServer.Add(findMovieByIdQuery.Name, findMovieByIdQuery)
}