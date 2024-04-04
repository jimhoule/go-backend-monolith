package repositories

import (
	"app/database/postgres"
	"app/profiles/domain/models"
	"context"
)

type PostgresProfilesRepository struct {
	Db *postgres.Db
}

func (ppr *PostgresProfilesRepository) FindAllByAccountId(accountId string) ([]*models.Profile, error) {
	query := "SELECT id, name, accountId FROM profiles WHERE accountId = $1"
	rows, err := ppr.Db.Query(context.Background(), query, accountId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	profiles := []*models.Profile{}
	for rows.Next() {
		profile := &models.Profile{}
		err = rows.Scan(&profile.Id, &profile.Name, &profile.AccountId)
		if err != nil {
			return nil, err
		}

		profiles = append(profiles, profile)
	}

	return profiles, nil
}

func (ppr *PostgresProfilesRepository) FindById(id string) (*models.Profile, error) {
	query := "SELECT id, name, accountId, languageId FROM profiles WHERE id = $1"
	row := ppr.Db.QueryRow(context.Background(), query, id)

	profile := &models.Profile{}
	err := row.Scan(&profile.Id, &profile.Name, &profile.AccountId, &profile.LanguageId)
	if err != nil {
		return nil, err
	}

	return profile, nil
}

func (ppr *PostgresProfilesRepository) Update(id string, profile *models.Profile) (*models.Profile, error) {
	query := "UPDATE profiles SET name = $1, languageId = $2 WHERE id = $3 RETURNING *"
	row := ppr.Db.QueryRow(context.Background(), query, profile.Name, profile.LanguageId, id)

	updatedProfile := &models.Profile{}
	err := row.Scan(&updatedProfile.Id, &updatedProfile.Name, &updatedProfile.AccountId, &updatedProfile.LanguageId)
	if err != nil {
		return nil, err
	}

	return updatedProfile, nil
}

func (ppr *PostgresProfilesRepository) Delete(id string) (string, error) {
	query := "DELETE FROM profiles WHERE id = $1"
	_, err := ppr.Db.Exec(context.Background(), query, id)
	if err != nil {
		return "", err
	}
	
	return id, nil
}

func (ppr *PostgresProfilesRepository) Create(profile *models.Profile) (*models.Profile, error) {
	query := "INSERT INTO profiles(id, name, accountId, languageId) VALUES(@id, @name, @accountId, @languageId)"
	args := postgres.NamedArgs{
		"id": profile.Id,
		"name": profile.Name,
		"accountId": profile.AccountId,
		"languageId": profile.LanguageId,
	}
	_, err := ppr.Db.Exec(context.Background(), query, args)
	if err != nil {
		return nil , err
	}

	return profile, nil
}