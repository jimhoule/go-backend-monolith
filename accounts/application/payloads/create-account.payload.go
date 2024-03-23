package payloads

type CreateAccountPayload struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
	PlanId    string
}