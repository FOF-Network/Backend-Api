package handler

import (
	"net/http"

	"github.com/FOF-Network/Backend-Api/db"
	"github.com/labstack/echo"
)

func Get(db db.DB) func(c echo.Context) error {
	return func(c echo.Context) error {
		id, err := db.GetIDFromToken(c.Request().Header.Get("authorization"))
		if err != nil || id == 0 {
			return c.JSON(http.StatusUnauthorized, nil)
		}

		
	}
}