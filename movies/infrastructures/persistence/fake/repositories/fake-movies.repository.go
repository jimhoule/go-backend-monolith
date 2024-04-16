package repositories

import (
	"app/movies/domain/models"
	"context"
	"fmt"
)

type FakeMoviesRepository struct {}

var movies []*models.Movie

func ResetFakeMoviesRepository() {
	movies = []*models.Movie{}
}

func (fmr *FakeMoviesRepository) FindAll() ([]*models.Movie, error) {
	return movies, nil
}

func (fmr *FakeMoviesRepository) FindById(id string) (*models.Movie, error) {
	for _, movie := range movies {
		if movie.Id == id {
			return movie, nil
		}
	}

	return nil, fmt.Errorf("the language with id %s does not exist", id)
}

func (fmr *FakeMoviesRepository) Update(ctx context.Context, id string, movie *models.Movie) (*models.Movie, error) {
	updatedMovie, err := fmr.FindById(id)
	if err != nil {
		return nil, err
	}

	updatedMovie.GenreId = movie.GenreId

	return updatedMovie, nil
}

func (fmr *FakeMoviesRepository) Delete(ctx context.Context, id string) (string, error) {
	for index, movie := range movies {
		if movie.Id == id {
			movies = append(movies[:index], movies[index + 1:]...)
			break
		}
	}

	return id, nil
}

func (fmr *FakeMoviesRepository) Create(ctx context.Context, movie *models.Movie) (*models.Movie, error) {
	movies = append(movies, movie)

	return movie, nil
}