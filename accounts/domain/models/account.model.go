package models

type Account struct {
	Id                    string `json:"id"`
	FirstName             string `json:"firstName"`
	LastName              string `json:"lastName"`
	Email                 string `json:"email"`
	Password              string `json:"password"`
	IsMembershipCancelled bool   `json:"isMembershipCancelled"`
	PlanId                string `json:"planId"`
}