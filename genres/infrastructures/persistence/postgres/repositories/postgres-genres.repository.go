package repositories

import (
	"app/database"
	"app/genres/domain/models"
	"context"
)

type PostgresGenresRepository struct {
	Db *database.Db
}

func (pgr *PostgresGenresRepository) FindAll() ([]*models.Genre, error) {
	query := "SELECT id FROM genres"
	rows, err := pgr.Db.Connection.Query(context.Background(), query)
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
	row := pgr.Db.Connection.QueryRow(context.Background(), query, id)

	genre := &models.Genre{}
	err := row.Scan(&genre.Id)
	if err != nil {
		return nil, err
	}

	return genre, nil
}

func (pgr *PostgresGenresRepository) Delete(ctx context.Context, id string) (string, error) {
	query := "DELETE FROM genres WHERE id = $1"
	_, err := pgr.Db.Connection.Exec(ctx, query, id)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (pgr *PostgresGenresRepository) Create(ctx context.Context, genre *models.Genre) (*models.Genre, error) {
	query := "INSERT INTO genres(id) VALUES(@id)"
	args := database.NamedArgs{
		"id": genre.Id,
	}

	_, err := pgr.Db.Connection.Exec(ctx, query, args)
	if err != nil {
		return nil, err
	}

	return genre, nil
}