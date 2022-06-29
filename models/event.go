package models

type Event struct {
	ID		uint	`json:"id" gorm:"primary_key"`
	Title	string	`json:"title"`
	Venue	string	`json:"venue"`
	Price	string	`json:"price"`
}

