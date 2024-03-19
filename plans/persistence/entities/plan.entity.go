package entities

type Plan struct {
	Id          string  `gorm:"not null"`
	Name        string  `gorm:"not null"`
	Description string  `gorm:"not null"`
	Price       float32 `gorm:"not null"`
}