package entities

type Account struct {
	//postgres.Model
	Id        string `gorm:"not null"`
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	Email     string `gorm:"not null"`
	Password  string `gorm:"not null"`
}