package handler

import (
	"encoding/json"
	"net/http"
)



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

func WeatherStackReq (name string, token string) (*WeatherStackRes, error) {
	Wres := new(WeatherStackRes)
	httpClient := http.Client{}

 	req, err := http.NewRequest("GET", "https://api.weatherstack.com/current", nil)
  	if err != nil {
    	return Wres, nil
  	}

  	q := req.URL.Query()
  	q.Add("access_key", token)
  	q.Add("query", name)
  	req.URL.RawQuery = q.Encode()

    res, err := httpClient.Do(req)
    if err != nil {
    	return Wres, err
    }
  	defer res.Body.Close()

  	json.NewDecoder(res.Body).Decode(Wres)

	return Wres, nil
}