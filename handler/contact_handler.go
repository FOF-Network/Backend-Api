package handler

import (
	"net/http"
	"Backend-Api/mydb"

	"github.com/labstack/echo"
)

func Get(db db.DB, env map[string]string) func(c echo.Context) error {
	return func(c echo.Context) error {
		id, err := db.GetIDFromToken(c.Request().Header.Get("authorization"))
		if err != nil || id == 0 {
			return c.JSON(http.StatusUnauthorized, nil)
		}
		
		
		csc := c.Param("csc")
		contacts, err := db.GetContacts(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, nil)
		}
		
		if csc == "true" {
			firstContacts := contacts
			db
		}


		for _, contact := range contacts {
			w, err := WeatherStackReq(contact.CityName, env["token"])
			if err != nil || id == 0 {
				return c.JSON(http.StatusUnauthorized, nil)
			}
			contact.GeoInfo = w
		}
			
		return c.JSON(http.StatusOK, contacts)	
	}
}