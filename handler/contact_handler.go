package handler

import (
	"Backend-Api/models"
	"Backend-Api/mydb"
	"log"
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
		
		user, err := db.GetUser(id)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, nil)
		}

		csc := c.QueryParam("csc")
		contacts, err := db.GetContacts(user.Cellphone)
		if err != nil {
			log.Print(err.Error())
			return c.JSON(http.StatusInternalServerError, nil)
		}
		
		if csc == "true" {
			var secondContacts []*models.ContactModel
			var HelpContacts []*models.ContactModel
			for _, contact := range contacts {
				cc, err := db.GetContacts(contact.Cellphone)
				if err != nil {
					return c.JSON(http.StatusInternalServerError, nil)
				}
				secondContacts = append(secondContacts, cc...)
			}
			contacts = secondContacts

			for _, cont := range contacts {
				oc, err := db.GetContactWithCell(cont.UserCellphone)
				if err != nil {
					return c.JSON(http.StatusInternalServerError, nil)
				}
				cont.Owners = []models.Owner{{FirstName: oc.FirstName, LastName: oc.LastName, Cellphone: oc.Cellphone}}
			}

			check := map[string]int{}
			andis := 0

			for _, cont := range contacts {
				if v, found := check[cont.Cellphone]; found {
					HelpContacts[v].Owners = append(HelpContacts[v].Owners, cont.Owners...)
				} else {
					HelpContacts = append(HelpContacts, cont)
					check[cont.Cellphone] = andis
					andis++
				}
			}
			
			contacts = HelpContacts
		} 


		for _, contact := range contacts {
			w, err := WeatherStackReq(contact.CityName, env["WHT_TOKEN"])
			if err != nil {
				contact.GeoInfo = nil
			}
			contact.GeoInfo = w
			contact.Cellphone = ""
		}
			
		return c.JSON(http.StatusOK, contacts)	
	}
}

func Add(db mydb.DB) func(c echo.Context) error {
	return func(c echo.Context) error {
		id, err := db.GetIDFromToken(c.Request().Header.Get("authorization"))
		if err != nil || id == 0 {
			log.Print(err.Error())
			return c.JSON(http.StatusUnauthorized, nil)
		}

		user, err := db.GetUser(id)
		if err != nil {
			log.Print(err.Error())
			return c.JSON(http.StatusUnauthorized, nil)
		}

		contact := new(models.ContactModel)
		err = c.Bind(contact)
		if err != nil {
			log.Print(err.Error())
			return c.JSON(http.StatusUnprocessableEntity, nil)
		}

		err = db.InsertContact(user.Cellphone, contact)
		if err != nil {
			log.Print(err.Error())
			return c.JSON(http.StatusServiceUnavailable, map[string]interface{}{"error": err.Error()})
		}

		return c.JSON(http.StatusOK, nil)
	}
}

func Edit(db mydb.DB) func(c echo.Context) error {
	return func(c echo.Context) error {
		id, err := db.GetIDFromToken(c.Request().Header.Get("authorization"))
		if err != nil || id == 0 {
			return c.JSON(http.StatusUnauthorized, nil)
		}
		
		user, err := db.GetUser(id)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, nil)
		}

		cidInt, _ := strconv.Atoi(c.QueryParam("id"))
		cUserCell, err := db.GetContactUserCell(uint(cidInt))
		if err != nil {
			return c.JSON(http.StatusServiceUnavailable, map[string]interface{}{"error": err.Error()})
		}
		
		if *cUserCell != user.Cellphone {
			log.Print(*cUserCell)
			log.Print(user.Cellphone)
			return c.JSON(http.StatusUnauthorized, nil)
		}

		contact := new(models.ContactModel)
		err = c.Bind(contact)
		if err != nil {
			return c.JSON(http.StatusUnprocessableEntity, nil)
		}

		err = db.UpdateContact(uint(cidInt), contact)

		if err != nil {
			return c.JSON(http.StatusServiceUnavailable, map[string]interface{}{"error": err.Error()})
		}
		return c.JSON(http.StatusOK, nil)
	}
}

func Delete(db mydb.DB) func(c echo.Context) error {
	return func(c echo.Context) error {
		id, err := db.GetIDFromToken(c.Request().Header.Get("authorization"))
		if err != nil || id == 0 {
			return c.JSON(http.StatusUnauthorized, nil)
		}

		user, err := db.GetUser(id)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, nil)
		}

		cid, err := strconv.Atoi(c.QueryParam("id"))
		if err != nil {
			return c.JSON(http.StatusUnprocessableEntity, nil)
		}

		uCell, err := db.GetContactUserCell(uint(cid))

		if err != nil {
			return c.JSON(http.StatusServiceUnavailable, nil)
		}


		if *uCell != user.Cellphone {
			return c.JSON(http.StatusUnauthorized, nil)
		}

		err = db.DeleteContact(uint(cid))

		if err != nil {
			return c.JSON(http.StatusServiceUnavailable, nil)
		}

		return c.JSON(http.StatusOK, nil)
	}
}