package payloads

type CreatePlanPayload struct {
	Name        string
	Description string
	Price       float32
}