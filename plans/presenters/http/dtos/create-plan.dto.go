package dtos

type CreatePlanDto struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `jons:"price"`
}