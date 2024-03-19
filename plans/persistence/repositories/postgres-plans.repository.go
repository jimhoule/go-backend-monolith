package repositories

import (
	"app/database/postgres"
	"app/plans/domain/models"
	"app/plans/persistence/entities"
	"app/plans/persistence/mappers"
	"log"
)

type PostgresPlansRepository struct {
	PlansMapper mappers.PlansMapper
	Db *postgres.Db
}

func (ppr *PostgresPlansRepository) FindAll() ([]*models.Plan, error) {
	var planEntities []*entities.Plan
	result := ppr.Db.Find(&planEntities)
	if result.Error != nil {
		return nil, result.Error
	}

	planModels := []*models.Plan{}
	for _, planEntity := range planEntities {
		planModel := ppr.PlansMapper.ToDomainModel(planEntity)
		planModels = append(planModels, planModel)
	}

	return planModels, nil
}

func (ppr *PostgresPlansRepository) FindById(id string) (*models.Plan, error) {
	var planEntities []*entities.Plan
	result := ppr.Db.Where("id = ?", id).Find(&planEntities)
	if result.Error != nil {
		log.Panicf("Error finding Plan with id %s: %s", id, result.Error.Error())
		return nil, result.Error
	}

	return ppr.PlansMapper.ToDomainModel(planEntities[0]), nil
}

func (ppr *PostgresPlansRepository) Create(planModel *models.Plan) (*models.Plan, error) {
	planEntity := ppr.PlansMapper.ToEntity(planModel)
	result := ppr.Db.Create(planEntity)
	if result.Error != nil {
		log.Panicf("Error saving a Plan: %s", result.Error.Error())
		return nil, result.Error
	}

	return planModel, nil
}