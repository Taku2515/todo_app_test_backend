package services

import (
	"github.com/labstack/echo/v4"

	"backend/crud"
	"net/http"
	"strconv"
)

func DeletePost(c echo.Context) error {

	type Body struct {
		UserName string `json:"user_name"`
	}

	idStr := c.Param("postId")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid post ID"})
	}

	obj := Body{}

	if err := c.Bind(&obj); err != nil {
		return err
	}

	_, err = crud.FetchPost(uint(id))

	if err != nil {
		return err
	}

	post := crud.RemovePost(uint(id))

	return c.JSON(http.StatusOK, post)

}
