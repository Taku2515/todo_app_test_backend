package services

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"backend/database"
	"backend/crud"
)

func NewUser(c echo.Context) error {


	user := database.User{}
	type body struct {
		Name string `json:"name"`
	}
	obj := body{}

	if err := c.Bind(&obj); err != nil {
		return err;
	}
	user.Name = obj.Name

	if user.Name == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "User name is required.")
	}

	user = crud.CreateUserDB(user)

	if user.ID == 0 {
		return echo.NewHTTPError(http.StatusConflict, "This username is already exist.")
	}

	return c.JSON(http.StatusCreated, user)


}