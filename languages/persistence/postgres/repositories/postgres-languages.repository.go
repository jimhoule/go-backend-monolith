package repositories

import (
	"app/database/postgres"
	"app/languages/domain/models"
	"context"
)

type PostgresLanguagesRepository struct {
	Db *postgres.Db
}

func (plr *PostgresLanguagesRepository) FindAll() ([]*models.Language, error) {
	query := "SELECT id, code, title FROM languages"
	rows, err := plr.Db.Connection.Query(context.Background(), query)
	if err != nil {
		return nil , err
	}
	defer rows.Close()

	languages := []*models.Language{}
	for rows.Next() {
		language := &models.Language{}
		err = rows.Scan(&language.Id, &language.Code, &language.Title)
		if err != nil {
			return nil, err
		}

		languages = append(languages, language)
	}

	return languages, nil
}

func (plr *PostgresLanguagesRepository) FindById(id string) (*models.Language, error) {
	query := "SELECT id, code, title FROM languages WHERE id = $1"
	row := plr.Db.Connection.QueryRow(context.Background(), query, id)


	language := &models.Language{}
	err := row.Scan(&language.Id, &language.Code, &language.Title)
	if err != nil {
		return nil, err
	}

	return language, nil
}

func (plr *PostgresLanguagesRepository) Update(id string, language *models.Language) (*models.Language, error) {
	query := "UPDATE languages SET code = $1, title = $2 WHERE id = $3 RETURNING *"
	row := plr.Db.Connection.QueryRow(context.Background(), query, language.Code, language.Title, id)


	updatedLanguage := &models.Language{}
	err := row.Scan(&updatedLanguage.Id, &updatedLanguage.Code, &updatedLanguage.Title)
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

func (plr *PostgresLanguagesRepository) Create(language *models.Language) (*models.Language, error) {
	query := "INSERT INTO languages(id, code, title) VALUES(@id, @code, @title)"
	args := postgres.NamedArgs{
		"id": language.Id,
		"code": language.Code,
		"title": language.Title,
	}
	_, err := plr.Db.Connection.Exec(context.Background(), query, args)
	if err != nil {
		return nil , err
	}

	return language, nil
}