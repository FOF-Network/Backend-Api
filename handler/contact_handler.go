package handler

import (
	"Backend-Api/mydb"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func Get(db mydb.DB, env map[string]string) func(c echo.Context) error {
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
			var secondContacts []*mydb.ContactModel
			for _, contact := range contacts {
				cc, err := db.GetContacts(contact.ID)
				if err != nil {
					return c.JSON(http.StatusInternalServerError, nil)
				}
				secondContacts = append(secondContacts, cc)
			}
			contacts = secondContacts
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
