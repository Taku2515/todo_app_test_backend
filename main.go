package main

import (

	"backend/database"
	"backend/services"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	database.Init()

	// POST
	e.POST("/create/user", services.NewUser)
	e.POST("/create/post", services.NewPost)

	// GET
	e.GET("/get/users", services.GetUsers)
	e.GET("/get/posts", services.GetPosts)

	// PUT
	e.PUT("update/username/:id", services.UpdateUser)

	// DELETE
	e.DELETE("delete/user/:id", services.DeleteUser)

	e.Logger.Fatal(e.Start(":8080"))
}