package services

import (
	"app/languages/application/payloads"
	"app/languages/application/ports"
	"app/languages/domain/factories"
	"app/languages/domain/models"
)

type LanguagesService struct {
	LanguagesFactory *factories.LanguagesFactory
	LanguagesRepository ports.LanguagesRepositoryPort
}

func (ls *LanguagesService) FindAll() ([]*models.Language, error) {
	return ls.LanguagesRepository.FindAll()
}

func (ls *LanguagesService) FindById(id string) (*models.Language, error) {
	return ls.LanguagesRepository.FindById(id)
}

func (ls *LanguagesService) Update(id string, updateLanguagePayload *payloads.UpdateLanguagePayload) (*models.Language, error) {
	return ls.LanguagesRepository.Update(id, &models.Language{
		Code: updateLanguagePayload.Code,
		Title: updateLanguagePayload.Title,
	})
}

func (ls *LanguagesService) Delete(id string) (string, error) {
	return ls.LanguagesRepository.Delete(id)
}

func (ls *LanguagesService) Create(createLanguagePayload *payloads.CreateLanguagePayload) (*models.Language, error) {
	language := ls.LanguagesFactory.Create(createLanguagePayload.Code, createLanguagePayload.Title)

	return ls.LanguagesRepository.Create(language)
}