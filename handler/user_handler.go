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

		err = db,
	}
}