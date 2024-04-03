package repositories

import (
	"app/languages/domain/models"
	"fmt"
)

var languages []*models.Language = []*models.Language{}

func ResetFakeLanguagesRepository() {
	languages = []*models.Language{}
}

type FakeLanguagesRepository struct {}

func (flr *FakeLanguagesRepository) FindAll() ([]*models.Language, error) {
	return languages, nil
}

func (flr *FakeLanguagesRepository) FindById(id string) (*models.Language, error) {
	for _, language := range languages {
		if language.Id == id {
			return language, nil
		}
	}

	return nil, fmt.Errorf("the language with id %s does not exist", id)
}

func (flr *FakeLanguagesRepository) Update(id string, language *models.Language) (*models.Language, error) {
	for _, languageToUpdate := range languages {
		if languageToUpdate.Id == id {
			languageToUpdate.Code = language.Code
			languageToUpdate.Title = language.Title
			return languageToUpdate, nil
		}
	}

	return nil, fmt.Errorf("the language with id %s does not exist", id)
}

func (flr *FakeLanguagesRepository) Delete(id string) (string, error) {
	for index, language := range languages {
		if language.Id == id {
			/*
			 * NOTES:
			 *	- profiles[:index] yields the slice elements before
			 *	- profiles[index+1:]... yields the slice elements after
			 *	- The 2 slices are then merged together by append()
			 */
			 languages = append(languages[:index], languages[index + 1:]...)
		}
	}

	return id, nil
}

func (flr *FakeLanguagesRepository) Create(language *models.Language) (*models.Language, error) {
	languages = append(languages, language)

	return language, nil
}