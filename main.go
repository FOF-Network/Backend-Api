package main

import (
	"time"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	env, err := godotenv.Read()
	if err != nil {
		logrus.Fatal(time.Now().Format(time.RFC1123))
	}


}