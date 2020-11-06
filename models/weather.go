package models

type WeatherStackRes struct {
	Success bool `json:"success"`
	Location struct {
		Name string `json:"string"`
		Country string `json:"country"`
		Localtime string `json:"localtime"`
	} `json:"location"`
	Current struct {
		Temperature int `json:"temperature"`
		WeatherDescriptions []string `json:"weather_descriptions"`

	} `json:"current"`
	Error   struct {
		Code int `json:"code"`
		Type string `json:"type"`
		Info string `json:info`
	} `json:"error"`
}