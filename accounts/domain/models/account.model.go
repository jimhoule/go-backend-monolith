package models

type Account struct {
	Id                    string
	FirstName             string
	LastName              string
	Email                 string
	Password              string
	IsMembershipCancelled bool
	PlanId                string
}