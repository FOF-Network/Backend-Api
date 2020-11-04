package db

import "time"

type contact struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	BirthDay  time.Time `json:"birth_day"`
	Email     string    `json:"email"`
	Job       string    `json:"job"`
	Interests string	`json:"interests"`
}

type DB interface {
	GetIDFromToken(token string) (uint, error)
	GetContacts(id uint) ([]*contact)
	SetToken(userID uint) error
}