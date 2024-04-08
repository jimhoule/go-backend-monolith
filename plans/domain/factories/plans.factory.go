package factories

import (
	"app/plans/domain/models"
	"app/uuid/services"
)

type PlansFactory struct{
	UuidService services.UuidService
}

func (pf *PlansFactory) Create(price float32) *models.Plan {
	return &models.Plan{
		Id: pf.UuidService.Generate(),
		Price: price,
	}
}