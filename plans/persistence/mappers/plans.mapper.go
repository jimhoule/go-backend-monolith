package mappers

import (
	"app/plans/domain/models"
	"app/plans/persistence/entities"
)

type PlansMapper struct{}

func (*PlansMapper) ToDomainModel(planEntity *entities.Plan) *models.Plan {
	return &models.Plan{
		Id: planEntity.Id,
		Name: planEntity.Name,
		Description: planEntity.Description,
		Price: planEntity.Price,
	}
}

func (*PlansMapper) ToEntity(planModel *models.Plan) *entities.Plan {
	return &entities.Plan{
		Id: planModel.Id,
		Name: planModel.Name,
		Description: planModel.Description,
		Price: planModel.Price,
	}
}