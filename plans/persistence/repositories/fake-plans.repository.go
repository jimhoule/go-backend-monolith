package repositories

import (
	"app/plans/domain/models"
	"app/plans/persistence/entities"
	"app/plans/persistence/mappers"
	"fmt"
)

type FakePlansRepository struct {
	PlansMapper mappers.PlansMapper
}

var planEntities []*entities.Plan = []*entities.Plan{}

func (fpr *FakePlansRepository) FindAll() ([]*models.Plan, error) {
	planModels := []*models.Plan{}

	for _, planEntity := range planEntities {
		planModel := fpr.PlansMapper.ToDomainModel(planEntity)
		planModels = append(planModels, planModel)
	}

	return planModels, nil
}

func (fpr *FakePlansRepository) FindById(id string) (*models.Plan, error) {
	for _, planEntity := range planEntities {
		if planEntity.Id == id {
			return fpr.PlansMapper.ToDomainModel(planEntity), nil
		}
	}

	return nil, fmt.Errorf("the plan with id %s does not exist", id)
}

func (fpr *FakePlansRepository) Create(planModel *models.Plan) (*models.Plan, error) {
	planEntity := fpr.PlansMapper.ToEntity(planModel)
	planEntities = append(planEntities, planEntity)

	return planModel, nil
}