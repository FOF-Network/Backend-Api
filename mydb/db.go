package mydb

import (
	"time"
	"Backend-Api/handler"
)

type ContactModel struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	BirthDay  time.Time `json:"birth_day"`
	Email     string    `json:"email"`
	Job       string    `json:"job"`
	Interests string	`json:"interests"`
	CityName    string    `json:"city_name"`
	GeoInfo    handler.WeatherStackRes `json:"geo_info"`
}

type DB interface {
	GetIDFromToken(token string) (uint, error)
	GetContacts(id uint) ([]*ContactModel, error)
	GetCity(id uint) ()
	SetToken(userID uint) error
}