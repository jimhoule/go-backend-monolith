package services

import (
	"app/translations/application/constants"
	"app/translations/application/payloads"
	"app/translations/domain/factories"
	"app/translations/domain/models"
	"app/translations/infrastructures/persistence/fake/repositories"
	"context"
	"testing"
)

func getTestContext() (*TranslationsService, func(), func() ([]*models.Translation, error)) {
	translationsService := &TranslationsService{
		TranslationsFactory: &factories.TranslationsFactory{},
		TranslationsRepository: &repositories.FakeTranslationsRepository{},
	}

	createTranslations := func() ([]*models.Translation, error) {
		return translationsService.CreateBatch(
			context.Background(),
			"fakeEntityId",
			constants.TanslationTypeLabel,
			[]*payloads.CreateTranslationPayload{
				{
					LanguageId: "fakeLanguageId",
					Text: "fakeText",
				},
			},
		)
	}
	
	return translationsService, repositories.ResetFakeTranslationRepository, createTranslations
}

func TestCreateTranslationBatchService(t *testing.T) {
	_, reset, createTranslations := getTestContext()
	defer reset()

	_, err := createTranslations()
	if err != nil {
		t.Errorf("Expected to create Translations but got %v", err)
	}
}

func TestFindAllTranslationsService(t *testing.T) {
	translationsService, reset, createTranslations := getTestContext()
	defer reset()

	newTranslations, _ := createTranslations()

	translations, err := translationsService.FindAll()
	if err != nil {
		t.Errorf("Expected slice of Translations but got %v", err)
		return
	}

	if len(translations) != len(newTranslations) {
		t.Errorf("Expected Translations slice id to be equal of equal size of New Tranlsations slice but got %d", len(translations))
	}
}

func TestFindAllTranslationsByEntityIdAndTypeService(t *testing.T) {
	translationsService, reset, createTranslations := getTestContext()
	defer reset()

	newTranslations, _ := createTranslations()

	translations, err := translationsService.FindAllByEntityIdAndType(newTranslations[0].EntityId, constants.TanslationTypeLabel)
	if err != nil {
		t.Errorf("Expected slice of Translations but got %v", err)
		return
	}

	if len(translations) != len(newTranslations) {
		t.Errorf("Expected Translations slice id to be equal of equal size of New Tranlsations slice but got %d", len(translations))
	}
}

func TestFindTranslationByComppsiteIdService(t *testing.T) {
	translationsService, reset, createTranslations := getTestContext()
	defer reset()

	newTranslations, _ := createTranslations()

	translation, err := translationsService.FindByCompositeId(
		newTranslations[0].EntityId,
		newTranslations[0].LanguageId,
		constants.TanslationTypeLabel,
	)
	if err != nil {
		t.Errorf("Expected slice of Translations but got %v", err)
		return
	}

	if translation.EntityId != newTranslations[0].EntityId {
		t.Errorf("Expected Translation entity id to be equal to New Translation entity id but got %s", translation.EntityId)
		return
	}

	if translation.LanguageId != newTranslations[0].LanguageId {
		t.Errorf("Expected Translation language id to be equal to New Translation language id but got %s", translation.LanguageId)
		return
	}

	if translation.Type != newTranslations[0].Type {
		t.Errorf("Expected Translation type to be equal to New Translation type but got %s", translation.Type)
	}
}

func TestUpsertTranslationBatchService(t *testing.T) {
	translationsService, reset, createTranslations := getTestContext()
	defer reset()

	newTranslations, _ := createTranslations()

	updatedText := "Updated fale translation text"
	translations, err := translationsService.UpsertBatch(
		context.Background(),
		newTranslations[0].EntityId,
		constants.TanslationTypeLabel,
		[]*payloads.UpdateTranslationPayload{
			{
				LanguageId: newTranslations[0].LanguageId,
				Text: updatedText,
			},
		},
	)
	if err != nil {
		t.Errorf("Expected slice of Translations but got %v", err)
		return
	}

	if len(translations) != len(newTranslations) {
		t.Errorf("Expected Translations slice id to be equal of equal size of New Tranlsations slice but got %d", len(translations))
	}

	if newTranslations[0].Text != updatedText {
		t.Errorf("Expected updated text to be equal to New Translation text but got %s", newTranslations[0].Text)
	}

	if translations[0].Text != updatedText {
		t.Errorf("Expected Translation text to be equal to updated text but got %s", translations[0].Text)
	}
}

func TestDeleteTranslationBatchService(t *testing.T) {
	translationsService, reset, createTranslations := getTestContext()
	defer reset()

	newTranslations, _ := createTranslations()

	entityId, err := translationsService.DeleteBatch(context.Background(), newTranslations[0].EntityId)
	if err != nil {
		t.Errorf("Expected entity id but got %v", err)
		return
	}

	if entityId != newTranslations[0].EntityId {
		t.Errorf("Expected entity id to be equal to New Translation entity id but got %s", entityId)
	}
}