package repositories

import (
	"app/accounts/domain/models"
	"app/database/postgres"
	"context"
)

type PostgresAccountsRepository struct{
	Db *postgres.Db
}

func (par *PostgresAccountsRepository) FindAll() ([]*models.Account, error) {
	query := "SELECT id, firstName, lastName, email, password, isMembershipCancelled, planId FROM accounts"
	rows, err := par.Db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	accounts := []*models.Account{}
	for rows.Next() {
		account := &models.Account{}

		err := rows.Scan(&account.Id, &account.FirstName, &account.LastName, &account.Email, &account.Password, &account.IsMembershipCancelled, &account.PlanId)
		if err != nil {
			return nil, err
		}

		accounts = append(accounts, account)
	}

	return accounts, nil
}

func (par *PostgresAccountsRepository) FindById(id string) (*models.Account, error) {
	query := "SELECT id, firstName, lastName, email, password, isMembershipCancelled, planId FROM accounts WHERE id = $1"
	row := par.Db.QueryRow(context.Background(), query, id)

	account := &models.Account{}
	err := row.Scan(&account.Id, &account.FirstName, &account.LastName, &account.Email, &account.Password, &account.IsMembershipCancelled, &account.PlanId)
	if err != nil {
		return nil, err
	}

	return account, nil
}

func (par *PostgresAccountsRepository) FindByEmail(email string) (*models.Account, error) {
	query := "SELECT id, firstName, lastName, email, password, isMembershipCancelled, planId FROM accounts WHERE email = $1"
	row := par.Db.QueryRow(context.Background(), query, email)

	account := &models.Account{}
	err := row.Scan(&account.Id, &account.FirstName, &account.LastName, &account.Email, &account.Password, &account.IsMembershipCancelled, &account.PlanId)
	if err != nil {
		return nil, err
	}

	return account, nil
}

func  (par *PostgresAccountsRepository) Create(account *models.Account) (*models.Account, error) {
	query := "INSERT INTO accounts (id, firstName, lastName, email, password, isMembershipCancelled, planId) VALUES (@id, @firstName, @lastName, @email, @password, @isMembershipCancelled, @planId)"
	args := postgres.NamedArgs{
		"id": account.Id,
		"firstName": account.FirstName,
		"lastName": account.LastName,
		"email": account.Email,
		"password": account.Password,
		"isMembershipCancelled": account.IsMembershipCancelled,
		"planId": account.PlanId,
	}
	_, err := par.Db.Exec(context.Background(), query, args)
	if err != nil {
		return nil, err
	}

	return account, nil
}