package main

import (
	"net/http"

	"backend/database"

	"github.com/labstack/echo/v4"
)

func connect(c echo.Context) error {
	db, _ := database.DB.DB()
	defer db.Close()
	err := db.Ping()
	if err != nil {
		return c.String(http.StatusInternalServerError, "DB接続失敗しました")
	} else {
		return c.String(http.StatusOK, "DB接続しました")
	}
}

func main() {
	e := echo.New()
	e.GET("/", connect)
	e.Logger.Fatal(e.Start(":8080"))
}