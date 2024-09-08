package services

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"backend/database"
	"backend/crud"
)

func GetUsers(c echo.Context) error {

	users := []database.User{}
	users = crud.FetchUsers(users)
	return c.JSON(http.StatusOK, users)

}