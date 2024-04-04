package services

import (
	"app/languages/application/payloads"
	"app/languages/domain/factories"
	"app/languages/domain/models"
	"app/languages/persistence/fake/repositories"
	"app/uuid"
	"testing"
)

func getTestContext() (*LanguagesService, func(), func() (*models.Language, error)) {
	languagesService := &LanguagesService{
		LanguagesFactory: &factories.LanguagesFactory{
			UuidService: uuid.GetService(),
		},
		LanguagesRepository: &repositories.FakeLanguagesRepository{},
	}

	createLanguage := func() (*models.Language, error) {
		return languagesService.Create(&payloads.CreateLanguagePayload{
			Code: "Fake code",
			Title: "Fake title",
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

	if languages[0] != newLanguage {
		t.Errorf("Expected first element of Languages slice to be equal to New Language but got %v", languages[0])
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

	if language != newLanguage {
		t.Errorf("Expected Language to equal New Language but got %v", language)
	}
}

func TestUpdateLanguageService(t *testing.T) {
	languagesService, reset, createLanguage := getTestContext()
	defer reset()

	newLanguage, _ := createLanguage()

	updatedTitle := "Updated fake language title"
	language, err := languagesService.Update(newLanguage.Id, &payloads.UpdateLanguagePayload{
		Code: newLanguage.Code,
		Title: updatedTitle,
	})
	if err != nil {
		t.Errorf("Expected Language but got %v", err)
		return
	}

	if newLanguage.Title != updatedTitle {
		t.Errorf("Expected New Language title to equal updated title but got %s", newLanguage.Title)
		return
	}

	if language.Title != updatedTitle {
		t.Errorf("Expected Language title to equal updated title but got %s", language.Title)
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

	if newLanguage.Id != languageId {
		t.Errorf("Expected New Language id to equal Language id but got %s", newLanguage.Id)
	}
}