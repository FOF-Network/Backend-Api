package main

import (
	"Backend-Api/handler"
	"Backend-Api/mydb"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

func main() {
	env, err := godotenv.Read()
	if err != nil {
		logrus.Fatal(err.Error())
	}

	ourDB, err := mydb.New(env)

	if err != nil {
		logrus.Fatal(err)
	}
	e := echo.New()

	e.POST("/register", handler.Register(ourDB))
	e.POST("/login", handler.LogIn(ourDB))
	e.PATCH("/user", handler.Update(ourDB))
	e.GET("/contact", handler.Get(ourDB, env))
	e.POST("/contact", handler.Add(ourDB))
	e.PATCH("/contact", handler.Edit(ourDB))
	e.DELETE("/contact", handler.Delete(ourDB))

	e.Logger.Fatal(e.Start(":8080"))
}