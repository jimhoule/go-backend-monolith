package repositories

import (
	"app/database/postgres"
	"app/translations/domain/models"
	"context"
)

type PostgresTranslationsRepository struct {
	Db *postgres.Db
}

func (ptr *PostgresTranslationsRepository) ExecuteTransaction(
	ctx context.Context,
	executeQuery func(ctx context.Context) (any, error),
) (any, error) {
	return postgres.ExecuteTransaction(ctx, executeQuery)
}

func (ptr *PostgresTranslationsRepository) FindAll() ([]*models.Translation, error) {
	query := "SELECT entity_id, language_id, text text FROM translations"
	rows, err := ptr.Db.Connection.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	translations := []*models.Translation{}
	for rows.Next() {
		translation := &models.Translation{}
		err = rows.Scan(&translation.EntityId, &translation.LanguageId, &translation.Text)
		if err != nil {
			return nil, err
		}

		translations = append(translations, translation)
	}

	return translations, nil
}

func (ptr *PostgresTranslationsRepository) FindAllByEntityId(entityId string) ([]*models.Translation, error) {
	query := "SELECT language_id, text FROM translations WHERE entity_id = $1"
	rows, err := ptr.Db.Connection.Query(context.Background(), query, entityId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	translations := []*models.Translation{}
	for rows.Next() {
		translation := &models.Translation{}
		err = rows.Scan(&translation.LanguageId, &translation.Text)
		if err != nil {
			return nil, err
		}

		translations = append(translations, translation)
	}

	return translations, nil
}

func (ptr *PostgresTranslationsRepository) FindByCompositeId(entityId string, languageId string) (*models.Translation, error) {
	query := "SELECT entity_id, language_id, text text FROM translations WHERE (entity_id, language_id) = ($1, $2)"
	row := ptr.Db.Connection.QueryRow(context.Background(), query, entityId, languageId)

	translation := &models.Translation{}
	err := row.Scan(&translation.EntityId, &translation.LanguageId, &translation.Text)
	if err != nil {
		return nil, err
	}

	return translation, nil
}

func (ptr *PostgresTranslationsRepository) UpdateBatch(ctx context.Context, translations []*models.Translation) ([]*models.Translation, error) {
	// Creates batch
	batch := &postgres.Batch{}
	for _, translation := range translations {
		query := "UPDATE translations SET text = $1 WHERE entity_id = $2 AND language_id = $3 RETURNING language_id, text"
		batch.Queue(query, translation.Text, translation.EntityId, translation.LanguageId)
	}

	// Executes batch
	batchResult := ptr.Db.Connection.SendBatch(ctx, batch)
	defer batchResult.Close()

	// Gets result of each batch update query
	updatedTranslations := []*models.Translation{}
	for range translations {
		row := batchResult.QueryRow()

		updatedTranslation := &models.Translation{}
		err := row.Scan(&updatedTranslation.LanguageId, &updatedTranslation.Text)
		if err != nil {
			return nil, err
		}

		updatedTranslations = append(updatedTranslations, updatedTranslation)
	}

	return updatedTranslations, nil
}

func (ptr *PostgresTranslationsRepository) UpsertBatch(ctx context.Context, translations []*models.Translation) ([]*models.Translation, error) {
	// Creates batch
	batch := &postgres.Batch{}
	for _, translation := range translations {
		query := `
			INSERT INTO translations(entity_id, language_id, text) VALUES($1, $2, $3)
			ON CONFLICT ON CONSTRAINT pk_translation DO UPDATE SET text = $3
			RETURNING language_id, text
		`
		batch.Queue(query, translation.EntityId, translation.LanguageId, translation.Text)
	}

	// Executes batch
	batchResult := ptr.Db.Connection.SendBatch(ctx, batch)
	defer batchResult.Close()

	// Gets result of each batch update query
	updatedTranslations := []*models.Translation{}
	for range translations {
		row := batchResult.QueryRow()

		updatedTranslation := &models.Translation{}
		err := row.Scan(&updatedTranslation.LanguageId, &updatedTranslation.Text)
		if err != nil {
			return nil, err
		}

		updatedTranslations = append(updatedTranslations, updatedTranslation)
	}

	return updatedTranslations, nil
}

func (ptr* PostgresTranslationsRepository) DeleteBatch(ctx context.Context, entityId string) (string, error) {
	query := "DELETE from translations WHERE entity_id = $1"
	_, err := ptr.Db.Connection.Exec(ctx, query, entityId)
	if err != nil {
		return "", err
	}

	return entityId, nil
}

func (ptr* PostgresTranslationsRepository) CreateBatch(ctx context.Context, translations []*models.Translation) ([]*models.Translation, error) {
	_, err := ptr.Db.Connection.CopyFrom(
		context.Background(),
		postgres.Identifier{"translations"},
		[]string{"entity_id", "language_id", "text"},
		ptr.Db.CopyFromSlice(len(translations), func(index int) ([]interface{}, error) {
			translation := translations[index]
			return []interface{}{translation.EntityId, translation.LanguageId, translation.Text}, nil
		}),
	)
	if err != nil {
		return nil, err
	}

	return translations, nil
}