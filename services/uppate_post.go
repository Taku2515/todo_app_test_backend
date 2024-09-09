package services

import (
	"github.com/labstack/echo/v4"

	"backend/crud"
	"net/http"
	"strconv"
)

func UpdatePost(c echo.Context) error {

	type Body struct {
		Task       string    `json:"task"`
		Importance uint      `json:"importance"`
		Deadline   string    `json:"deadline"`
	}

	idStr := c.Param("postId")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
	}

	obj := Body{}

	if err := c.Bind(&obj); err != nil {
		return err
	}

	post, err := crud.FetchPost(uint(id))

	if err != nil {
		return err
	}

	post = crud.ChangePost(post, obj.Task, obj.Importance, obj.Deadline)

	return c.JSON(http.StatusOK, post)

}
