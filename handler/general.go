package handler

import (
	"encoding/json"
	"net/http"
	"Backend-Api/models"
)


func WeatherStackReq (name string, token string) (*models.WeatherStackRes, error) {
	Wres := new(models.WeatherStackRes)
	httpClient := http.Client{}

 	req, err := http.NewRequest("GET", "http://api.weatherstack.com/current", nil)
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