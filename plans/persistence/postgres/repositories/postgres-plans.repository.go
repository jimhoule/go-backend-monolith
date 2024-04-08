package repositories

import (
	"app/database/postgres"
	"app/plans/domain/models"
	"context"
)

type PostgresPlansRepository struct {
	Db *postgres.Db
}

func (ppr *PostgresPlansRepository) FindAll() ([]*models.Plan, error) {
	query := "SELECT id, price FROM plans"
	rows, err := ppr.Db.Connection.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	plans := []*models.Plan{}
	for rows.Next() {
		plan := &models.Plan{}

		err := rows.Scan(&plan.Id, &plan.Price)
		if err != nil {
			return nil, err
		}

		plans = append(plans, plan)
	}

	return plans, nil
}

func (ppr *PostgresPlansRepository) FindById(id string) (*models.Plan, error) {
	query := "SELECT id, price FROM plans WHERE id = $1"
	row := ppr.Db.Connection.QueryRow(context.Background(), query, id)

	plan := &models.Plan{}
	err := row.Scan(&plan.Id, &plan.Price)
	if err != nil {
		return nil, err
	}

	return plan, nil
}

func (ppr *PostgresPlansRepository) Create(ctx context.Context, plan *models.Plan) (*models.Plan, error) {
	query := "INSERT INTO plans (id, price) VALUES ($1, $2)"
	_, err := ppr.Db.Connection.Exec(ctx, query, plan.Id, plan.Price)
	if err != nil {
		return nil, err
	}

	return plan, nil
}