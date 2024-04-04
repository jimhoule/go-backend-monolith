package translations

import (
	"app/database/postgres"
	"app/translations/application/services"
	"app/translations/domain/factories"
	"app/translations/persistence/postgres/repositories"
)

func GetService(db *postgres.Db) *services.TranslationsService {
	return &services.TranslationsService{
		TranslationsFactory: &factories.TranslationsFactory{},
		TranslationsRepository: &repositories.PostgresTranslationsRepository{
			Db: db,
		},
	}
}