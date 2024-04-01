package dtos

type CreateProfileDto struct {
	Name      string `json:"name"`
	AccountId string `json:"accountId"`
}