package services

import (
	"github.com/labstack/echo/v4"

	"backend/crud"
	"net/http"
	"strconv"
)

func DeleteUser(c echo.Context) error {

	type Body struct {
		Name string `json:"name"`
	}

	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
	}

	obj := Body{}

	if err := c.Bind(&obj); err != nil {
		return err
	}

	_, err = crud.FetchUser(uint(id))

	if err != nil {
		return err
	}

	user := crud.RemoveUser(uint(id))

	return c.JSON(http.StatusOK, user)

}
