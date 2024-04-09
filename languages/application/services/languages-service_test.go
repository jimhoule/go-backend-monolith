package services

import (
	"app/languages/application/payloads"
	languagesFactories "app/languages/domain/factories"
	"app/languages/domain/models"
	"app/languages/persistence/fake/repositories"
	languagesRepositories "app/languages/persistence/fake/repositories"
	transactionsServices "app/transactions/application/services"
	transactionsRepositories "app/transactions/persistence/fake/repositories"
	translationsServices "app/translations/application/services"
	translationsFactories "app/translations/domain/factories"
	translationsRepositories "app/translations/persistence/fake/repositories"
	"app/uuid"
	"testing"
)

func getTestContext() (*LanguagesService, func(), func() (*models.Language, error)) {
	languagesService := &LanguagesService{
		LanguagesFactory: &languagesFactories.LanguagesFactory{
			UuidService: uuid.GetService(),
		},
		LanguagesRepository: &languagesRepositories.FakeLanguagesRepository{},
		TranslationsService: &translationsServices.TranslationsService{
			TranslationsFactory: &translationsFactories.TranslationsFactory{},
			TranslationsRepository: &translationsRepositories.FakeTranslationsRepository{},
		},
		TransactionsService: &transactionsServices.TransactionsService{
			TransactionsRepository: &transactionsRepositories.FakeTransactionsRepository{},
		},
	}

	createLanguage := func() (*models.Language, error) {
		return languagesService.Create(&payloads.CreateLanguagePayload{
			Code: "Fake code",
		})
	}

	return languagesService, repositories.ResetFakeLanguagesRepository, createLanguage
}

func TestCreateLanguageService(t *testing.T) {
	_, reset, createLanguage := getTestContext()
	defer reset()

	_, err := createLanguage()
	if err != nil {
		t.Errorf("Expected to create Language but got %v", err)
	}
}

func TestFindAllLanguagesService(t *testing.T) {
	languagesService, reset, createLanguage := getTestContext()
	defer reset()

	newLanguage, _ := createLanguage()

	languages, err := languagesService.FindAll()
	if err != nil {
		t.Errorf("Expected slice of Languages but got %v", err)
		return
	}

	if languages[0].Id != newLanguage.Id {
		t.Errorf("Expected Language id to be equal to New Language id but got %v", languages[0].Id)
	}
}

func TestFindLanguageByIdService(t *testing.T) {
	languagesService, reset, createLanguage := getTestContext()
	defer reset()

	newLanguage, _ := createLanguage()

	language, err := languagesService.FindById(newLanguage.Id)
	if err != nil {
		t.Errorf("Expected Language but got %v", err)
		return
	}

	if language.Id != newLanguage.Id {
		t.Errorf("Expected Language to equal New Language but got %v", language.Id)
	}
}

func TestUpdateLanguageService(t *testing.T) {
	languagesService, reset, createLanguage := getTestContext()
	defer reset()

	newLanguage, _ := createLanguage()

	updatedCode := "Updated fake language code"
	language, err := languagesService.Update(newLanguage.Id, &payloads.UpdateLanguagePayload{
		Code: updatedCode,
	})
	if err != nil {
		t.Errorf("Expected Language but got %v", err)
		return
	}

	if updatedCode != newLanguage.Code {
		t.Errorf("Expected updated code to equal New Language code but got %s", updatedCode)
		return
	}

	if updatedCode != language.Code {
		t.Errorf("Expected updated code to equal Language code but got %s", updatedCode)
	}
}

func TestDeleteLanguageService(t *testing.T) {
	languagesService, reset, createLanguage := getTestContext()
	defer reset()

	newLanguage, _ := createLanguage()

	languageId, err := languagesService.Delete(newLanguage.Id)
	if err != nil {
		t.Errorf("Expected Language id but got %v", err)
		return
	}

	if languageId != newLanguage.Id {
		t.Errorf("Expected Language id to equal New Language id but got %s", languageId)
	}
}