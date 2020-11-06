package models

import "time"

type User struct {
	ID        uint      `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Cellphone string    `json:"cellphone"`
	BirthDay  time.Time `json:"birth_day"`
	Password  string    `json:"-"`
	Email     string    `json:"email"`
	Job       string    `json:"job"`
	Interests string	`json:"interests"`
	CityName    string    `json:"city_name"`
	GeoInfo    WeatherStackRes `json:"geo_info"`
}
