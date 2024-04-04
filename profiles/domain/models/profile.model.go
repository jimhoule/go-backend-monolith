package models

type Profile struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	AccountId  string `json:"accountId,omitempty"`
	LanguageId string `json:"languageId,omitempty"`
}