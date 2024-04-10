package repositories

import (
	"app/genres/domain/models"
	"context"
	"fmt"
)

var genres []*models.Genre = []*models.Genre{}

func ResetFakeGenresRepository() {
	genres = []*models.Genre{}
}

type FakeGenresRepository struct {}

func (fgr *FakeGenresRepository) FindAll() ([]*models.Genre, error) {
	return genres, nil
}

func (fgr *FakeGenresRepository) FindById(id string) (*models.Genre, error) {
	for _, genre := range genres {
		if genre.Id == id {
			return genre, nil
		}
	}

	return nil, fmt.Errorf("the genre with id %s does not exist", id)
}

func (fgr *FakeGenresRepository) Delete(ctx context.Context, id string) (string, error) {
	for index, genre := range genres {
		if genre.Id == id {
			/*
			 * NOTES:
			 *	- genres[:index] yields the slice elements before
			 *	- genres[index + 1:]... yields the slice elements after
			 *	- The 2 slices are then merged together by append()
			 */
			 genres = append(genres[:index], genres[index + 1:]...)
			 break
		}
	}

	return id, nil
}

func (fgr *FakeGenresRepository) Create(ctx context.Context, genre *models.Genre) (*models.Genre, error) {
	genres = append(genres, genre)

	return genre, nil
}