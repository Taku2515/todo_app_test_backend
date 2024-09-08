package main

import (

	"backend/database"
	"backend/services"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	database.Init()

	e.POST("/create/user", services.NewUser)

	e.Logger.Fatal(e.Start(":8080"))
}