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
	rows, err := ptr.Db.Connection.Query(context.Background(), query)
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

func (ptr *PostgresTranslationsRepository) FindAllByEntityId(entityId string) ([]*models.Translation, error) {
	query := "SELECT languageCode, text FROM translations WHERE entityId = $1"
	rows, err := ptr.Db.Connection.Query(context.Background(), query, entityId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	translations := []*models.Translation{}
	for rows.Next() {
		translation := &models.Translation{}
		err = rows.Scan(&translation.LanguageCode, &translation.Text)
		if err != nil {
			return nil, err
		}

		translations = append(translations, translation)
	}

	return translations, nil
}

func (ptr *PostgresTranslationsRepository) FindByCompositeId(entityId string, languageCode string) (*models.Translation, error) {
	query := "SELECT entityId, languageCode, text FROM translations WHERE (entityId, languageCode) = ($1, $2)"
	row := ptr.Db.Connection.QueryRow(context.Background(), query, entityId, languageCode)

	translation := &models.Translation{}
	err := row.Scan(&translation.EntityId, &translation.LanguageCode, &translation.Text)
	if err != nil {
		return nil, err
	}

	return translation, nil
}

func (ptr* PostgresTranslationsRepository) Create(translations []*models.Translation) ([]*models.Translation, error) {
	_, err := ptr.Db.Connection.CopyFrom(
		context.Background(),
		postgres.Identifier{"translations"},
		[]string{"entityid", "languagecode", "text"},
		ptr.Db.CopyFromSlice(len(translations), func(index int) ([]interface{}, error) {
			translation := translations[index]
			return []interface{}{translation.EntityId, translation.LanguageCode, translation.Text}, nil
		}),
	)
	if err != nil {
		return nil, err
	}

	return translations, nil
}