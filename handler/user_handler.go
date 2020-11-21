package handler

import (
	"Backend-Api/models"
	"Backend-Api/mydb"
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/thanhpk/randstr"
)

func Register(db mydb.DB) func(c echo.Context) error {
	return func(c echo.Context) error {
		user := new(models.User)
		err := c.Bind(user)
		if err != nil {
			return c.JSON(http.StatusUnprocessableEntity, nil)
		}
		u, err := db.GetUserWithCellphone(user.Cellphone)

		if u != nil {
			log.Print(err.Error())
			return c.JSON(http.StatusServiceUnavailable, nil)
		}

		err = db.InsertUser(user)
		if err != nil {
			log.Print(err.Error())
			return c.JSON(http.StatusServiceUnavailable, nil)
		}

		return c.JSON(http.StatusOK, nil)
	}
}

func Update(db mydb.DB) func(c echo.Context) error {
	return func(c echo.Context) error {
		id, err := db.GetIDFromToken(c.Request().Header.Get("authorization"))
		if err != nil || id == 0 {
			return c.JSON(http.StatusUnauthorized, nil)
		}

		user := new(models.User)
		err = c.Bind(user)
		if err != nil {
			return c.JSON(http.StatusUnprocessableEntity, nil)
		}

		err = db.UpdateUser(id, user)
		if err != nil {
			return c.JSON(http.StatusServiceUnavailable, nil)
		}

		return c.JSON(http.StatusOK, nil)
	}
}

func LogIn(db mydb.DB) func(c echo.Context) error {
	return func(c echo.Context) error {
		user := new(models.User)
		err := c.Bind(user)
		if err != nil {
			return c.JSON(http.StatusUnprocessableEntity, nil)
		}

		oUser, err := db.GetUserWithCellphone(user.Cellphone)
		if err != nil {
			log.Print(err.Error())
			return c.JSON(http.StatusServiceUnavailable, nil)
		}

		if oUser.Password == user.Password {
			token := randstr.String(10) 
		
			err = db.SetToken(oUser.ID, token)
			if err != nil {
				log.Print(err.Error())
				return c.JSON(http.StatusServiceUnavailable, nil)
			}
			return c.JSON(http.StatusOK, echo.Map{"token":token})
		}
		return c.JSON(http.StatusUnauthorized, nil)
	}
}
