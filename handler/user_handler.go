package handler

import (
	"Backend-Api/models"
	"Backend-Api/mydb"
	"net/http"

	"github.com/labstack/echo"
)

func Register(db mydb.DB) func(c echo.Context) error {
	return func(c echo.Context) error {
		user := new(models.User)
		err := c.Bind(user)
		if err != nil {
			return c.JSON(http.StatusUnprocessableEntity, nil)
		}

		err = db.InsertUser(user)
		if err != nil {
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

		oUser, err := db.GetUser(user.ID)
		if err != nil {
			return c.JSON(http.StatusServiceUnavailable, nil)
		}

		if oUser.Password == user.Password {
			token := "lmvfdmvps"
			err = db.SetToken(oUser.ID, token)
			if err != nil {
				return c.JSON(http.StatusServiceUnavailable, nil)
			}
			return c.JSON(http.StatusOK, nil)
		}
		return c.JSON(http.StatusUnauthorized, nil)
	}
}
