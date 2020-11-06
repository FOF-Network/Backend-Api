package models

import "time"

type ContactModel struct {
	ID        uint      `json:"-"`
	UserID    uint      `json:"user_id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	BirthDay  time.Time `json:"birth_day"`
	Email     string    `json:"email"`
	Job       string    `json:"job"`
	Interests string	`json:"interests"`
	CityName    string    `json:"city_name"`
	GeoInfo    WeatherStackRes `json:"geo_info"`
}
