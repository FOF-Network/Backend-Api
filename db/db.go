package db

import "time"

type Contact struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	BirthDay  time.Time `json:"birth_day"`
	Email     string    `json:"email"`
	Job       string    `json:"job"`
	Interests string	`json:"interests"`
	CityName    uint    `json:"city_name"`
	CityTimeAndWaether 
}

type DB interface {
	GetIDFromToken(token string) (uint, error)
	GetContacts(id uint) ([]*Contact, error)
	GetCity(id uint) ()
	SetToken(userID uint) error
}