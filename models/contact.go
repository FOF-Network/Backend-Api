package models

type ContactModel struct {
	ID         uint            `json:"id"`
	UserCellphone    string       `json:"user_cellphone"`
	FirstName  string          `json:"first_name"`
	LastName   string          `json:"last_name"`
	Cellphone  string    	   `json:"cellphone"`
	BirthDay   int 			   `json:"birth_day"`
	BirthMonth int 			   `json:"birth_month"`
	Email      string    	   `json:"email"`
	Job        string          `json:"job"`
	Interests  string		   `json:"interests"`
	CityName   string    	   `json:"city_name"`
	GeoInfo    *WeatherStackRes `json:"geo_info"`
}
