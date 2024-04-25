package repositories

import (
	"app/database"
	"app/languages/domain/models"
	"context"
)

type PostgresLanguagesRepository struct {
	Db *database.Db
}

func (plr *PostgresLanguagesRepository) FindAll() ([]*models.Language, error) {
	query := "SELECT id, code FROM languages"
	rows, err := plr.Db.Connection.Query(context.Background(), query)
	if err != nil {
		return nil , err
	}
	defer rows.Close()

	languages := []*models.Language{}
	for rows.Next() {
		language := &models.Language{}
		err = rows.Scan(&language.Id, &language.Code)
		if err != nil {
			return nil, err
		}

		languages = append(languages, language)
	}

	return languages, nil
}

func (plr *PostgresLanguagesRepository) FindById(id string) (*models.Language, error) {
	query := "SELECT id, code FROM languages WHERE id = $1"
	row := plr.Db.Connection.QueryRow(context.Background(), query, id)


	language := &models.Language{}
	err := row.Scan(&language.Id, &language.Code)
	if err != nil {
		return nil, err
	}

	return language, nil
}

func (plr *PostgresLanguagesRepository) Update(ctx context.Context, id string, language *models.Language) (*models.Language, error) {
	query := "UPDATE languages SET code = $1 WHERE id = $2 RETURNING *"
	row := plr.Db.Connection.QueryRow(ctx, query, language.Code, id)


	updatedLanguage := &models.Language{}
	err := row.Scan(&updatedLanguage.Id, &updatedLanguage.Code)
	if err != nil {
		return nil, err
	}

	return updatedLanguage, nil
}

func (plr *PostgresLanguagesRepository) Delete(id string) (string, error) {
	query := "DELETE FROM languages WHERE id = $1"
	_, err := plr.Db.Connection.Exec(context.Background(), query, id)
	if err != nil {
		return "", err
	}
	
	return id, nil
}

func (plr *PostgresLanguagesRepository) Create(ctx context.Context, language *models.Language) (*models.Language, error) {
	query := "INSERT INTO languages(id, code) VALUES(@id, @code)"
	args := database.NamedArgs{
		"id": language.Id,
		"code": language.Code,
	}
	_, err := plr.Db.Connection.Exec(ctx, query, args)
	if err != nil {
		return nil , err
	}

	return language, nil
}