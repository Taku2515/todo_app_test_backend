package services

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"backend/database"
	"backend/crud"
)

func GetPosts(c echo.Context) error {

	posts := []database.Post{}
	posts = crud.FetchPosts(posts)
	return c.JSON(http.StatusOK, posts)

}