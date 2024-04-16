package factories

import (
	"app/movies/domain/models"
	"app/uuid/services"
)

type MoviesFactory struct {
	UuidService services.UuidService
}

func (mf *MoviesFactory) Create(genreId string) *models.Movie {
	return &models.Movie{
		Id: mf.UuidService.Generate(),
		GenreId: genreId,
	}
}