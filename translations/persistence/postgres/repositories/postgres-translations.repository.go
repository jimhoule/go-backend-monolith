package repositories

import (
	"app/database/postgres"
	"app/translations/domain/models"
	"context"
)

type PostgresTranslationsRepository struct {
	Db *postgres.Db
}

func (ptr *PostgresTranslationsRepository) FindAll() ([]*models.Translation, error) {
	query := "SELECT entityId, languageCode, text FROM translations"
	rows, err := ptr.Db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	translations := []*models.Translation{}
	for rows.Next() {
		translation := &models.Translation{}
		err = rows.Scan(&translation.EntityId, &translation.LanguageCode, &translation.Text)
		if err != nil {
			return nil, err
		}

		translations = append(translations, translation)
	}

	return translations, nil
}

func (ptr *PostgresTranslationsRepository) FindByCompositeId(entityId string, languageCode string) (*models.Translation, error) {
	query := "SELECT entityId, languageCode, text FROM translations WHERE (entityId, languageCode) = ($1, $2)"
	row := ptr.Db.QueryRow(context.Background(), query, entityId, languageCode)

	translation := &models.Translation{}
	err := row.Scan(&translation.EntityId, &translation.LanguageCode, &translation.Text)
	if err != nil {
		return nil, err
	}

	return translation, nil
}

func (ptr* PostgresTranslationsRepository) Create(translation *models.Translation) (*models.Translation, error) {
	query := "INSERT INTO translations(entityId, languageCode, text) VALUES(@entityId, @languageCode, @text)"
	args := postgres.NamedArgs{
		"entityId": translation.EntityId,
		"languageCode": translation.LanguageCode,
		"text": translation.Text,
	}
	_, err := ptr.Db.Exec(context.Background(), query, args)
	if err != nil {
		return nil, err
	}

	return translation, nil
}