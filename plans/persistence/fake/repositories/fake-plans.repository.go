package repositories

import (
	"app/plans/domain/models"
	"fmt"

	_ "app/plans/application/ports"
)

type FakePlansRepository struct{}

var plans []*models.Plan = []*models.Plan{}

func (fpr *FakePlansRepository) FindAll() ([]*models.Plan, error) {
	return plans, nil
}

func (fpr *FakePlansRepository) FindById(id string) (*models.Plan, error) {
	for _, plan := range plans {
		if plan.Id == id {
			return plan, nil
		}
	}

	return nil, fmt.Errorf("the account with id %s does not exist", id)
}

func (fpr *FakePlansRepository) Create(plan *models.Plan) (*models.Plan, error) {
	plans = append(plans, plan);

	return plan, nil
}