package factories

import (
	"app/genres/domain/models"
	"app/uuid/services"
)

type GenresFactory struct {
	UuidService services.UuidService
}

func (gf *GenresFactory) Create() *models.Genre {
	return &models.Genre{
		Id: gf.UuidService.Generate(),
	}
}