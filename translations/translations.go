package translations

import (
	"app/database"
	"app/translations/application/services"
	"app/translations/domain/factories"
	"app/translations/infrastructures/persistence/postgres/repositories"
)

func GetService(db *database.Db) *services.TranslationsService {
	return &services.TranslationsService{
		TranslationsFactory: &factories.TranslationsFactory{},
		TranslationsRepository: &repositories.PostgresTranslationsRepository{
			Db: db,
		},
	}
}