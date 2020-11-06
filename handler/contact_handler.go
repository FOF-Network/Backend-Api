package handler

import (
	"Backend-Api/models"
	"Backend-Api/mydb"
	"net/http"
	"strconv"

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
			var secondContacts []*models.ContactModel
			for _, contact := range contacts {
				cc, err := db.GetContacts(contact.ID)
				if err != nil {
					return c.JSON(http.StatusInternalServerError, nil)
				}
				secondContacts = append(secondContacts, cc...)
			}
			contacts = secondContacts
		} 


		for _, contact := range contacts {
			w, err := WeatherStackReq(contact.CityName, env["token"])
			if err != nil || id == 0 {
				return c.JSON(http.StatusUnauthorized, nil)
			}
			contact.GeoInfo = *w
		}
			
		return c.JSON(http.StatusOK, contacts)	
	}
}

func Add(db mydb.DB) func(c echo.Context) error {
	return func(c echo.Context) error {
		id, err := db.GetIDFromToken(c.Request().Header.Get("authorization"))
		if err != nil || id == 0 {
			return c.JSON(http.StatusUnauthorized, nil)
		}

		contact := new(models.ContactModel)
		err = c.Bind(contact)
		if err != nil {
			return c.JSON(http.StatusUnprocessableEntity, nil)
		}

		err = db.InsertContact(id, contact)
		if err != nil {
			return c.JSON(http.StatusOK, nil)
		}

		return c.JSON(http.StatusServiceUnavailable, map[string]interface{}{"error": err.Error()})
	}
}

func Edit(db mydb.DB) func(c echo.Context) error {
	return func(c echo.Context) error {
		id, err := db.GetIDFromToken(c.Request().Header.Get("authorization"))
		if err != nil || id == 0 {
			return c.JSON(http.StatusUnauthorized, nil)
		}

		contact := new(models.ContactModel)
		err = c.Bind(contact)
		if err != nil {
			return c.JSON(http.StatusUnprocessableEntity, nil)
		}

		err = db.UpdateContact(id, contact)
		if err != nil {
			return c.JSON(http.StatusOK, nil)
		}

		return c.JSON(http.StatusServiceUnavailable, map[string]interface{}{"error": err.Error()})
	}
}

func Delete(db mydb.DB) func(c echo.Context) error {
	return func(c echo.Context) error {
		id, err := db.GetIDFromToken(c.Request().Header.Get("authorization"))
		if err != nil || id == 0 {
			return c.JSON(http.StatusUnauthorized, nil)
		}

		cid, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusUnprocessableEntity, nil)
		}

		contact, err := db.GetContact(uint(cid))

		if err != nil {
			return c.JSON(http.StatusServiceUnavailable, nil)
		}

		if contact.UserID != id {
			return c.JSON(http.StatusUnauthorized, nil)
		}

		err = db.DeleteContact(uint(cid))

		if err != nil {
			return c.JSON(http.StatusServiceUnavailable, nil)
		}

		return c.JSON(http.StatusOK, nil)
	}
}