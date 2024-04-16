package repositories

import (
	"app/database/postgres"
	"app/movies/domain/models"
	"context"
)

type PostgresMoviesRepository struct {
	Db *postgres.Db
}

func (pmr *PostgresMoviesRepository) FindAll() ([]*models.Movie, error) {
	query := "SELECT id, genre_id FROM movies"
	rows, err := pmr.Db.Connection.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	movies := []*models.Movie{}
	for rows.Next() {
		movie := &models.Movie{}
		err = rows.Scan(&movie.Id, &movie.GenreId)
		if err != nil {
			return nil, err
		}

		movies = append(movies, movie)
	}

	return movies, nil
}

func (pmr *PostgresMoviesRepository) FindById(id string) (*models.Movie, error) {
	query := "SELECT id, genre_id FROM movies WHERE id = $1"
	row := pmr.Db.Connection.QueryRow(context.Background(), query, id)

	movie := &models.Movie{}
	err := row.Scan(&movie.Id, &movie.GenreId)
	if err != nil {
		return nil, err
	}

	return movie, nil
}

func (pmr *PostgresMoviesRepository) Update(ctx context.Context, id string, movie *models.Movie) (*models.Movie, error) {
	query := "UPDATE movies SET genre_id = $1 WHERE id = $2 RETURNING id, genre_id"
	row := pmr.Db.Connection.QueryRow(ctx, query, movie.GenreId, id)

	updatedMovie := &models.Movie{}
	err := row.Scan(&updatedMovie.Id, &updatedMovie.GenreId)
	if err != nil {
		return nil, err
	}

	return updatedMovie, nil
}

func (pmr *PostgresMoviesRepository) Delete(ctx context.Context, id string) (string, error) {
	query := "DELETE FROM movies WHERE id = $1"
	_, err := pmr.Db.Connection.Exec(ctx, query, id)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (pmr *PostgresMoviesRepository) Create(ctx context.Context, movie *models.Movie) (*models.Movie, error) {
	query := "INSERT INTO movies(id, genre_id) VALUES($1, $2)"
	row := pmr.Db.Connection.QueryRow(ctx, query, movie.Id, movie.GenreId)

	newMovie := &models.Movie{}
	err := row.Scan(&newMovie.Id, &newMovie.GenreId)
	if err != nil {
		return nil, err
	}

	return newMovie, nil
}