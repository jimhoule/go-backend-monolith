package repositories

import (
	"app/database/postgres"
	"app/genres/domain/models"
	"context"
)

type PostgresGenresRepository struct {
	Db *postgres.Db
}

func (pgr *PostgresGenresRepository) FindAll() ([]*models.Genre, error) {
	query := "SELECT id FROM genres"
	rows, err := pgr.Db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	genres := []*models.Genre{}
	for rows.Next() {
		genre := &models.Genre{}
		err = rows.Scan(&genre.Id)
		if err != nil {
			return nil, err
		}

		genres = append(genres, genre)
	}

	return genres, nil
}

func (pgr *PostgresGenresRepository) FindById(id string) (*models.Genre, error) {
	query := "SELECT id FROM genres WHERE id = $1"
	row := pgr.Db.QueryRow(context.Background(), query, id)

	genre := &models.Genre{}
	err := row.Scan(&genre.Id)
	if err != nil {
		return nil, err
	}

	return genre, nil
}

func (pgr *PostgresGenresRepository) Create(genre *models.Genre) (*models.Genre, error) {
	query := "INSERT INTO genres(id) VALUES(@id)"
	args := postgres.NamedArgs{
		"id": genre.Id,
	}

	_, err := pgr.Db.Exec(context.Background(), query, args)
	if err != nil {
		return nil, err
	}

	return genre, nil
}