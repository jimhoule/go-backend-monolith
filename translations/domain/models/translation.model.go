package models

type Translation struct {
	EntityId     string `json:"entityId,omitempty"`
	LanguageCode string `json:"languageCode,omitempty"`
	Text         string `json:"text"`
}