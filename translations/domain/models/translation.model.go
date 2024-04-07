package models

type Translation struct {
	EntityId   string `json:"entityId,omitempty"`
	LanguageId string `json:"languageId,omitempty"`
	Text       string `json:"text"`
}