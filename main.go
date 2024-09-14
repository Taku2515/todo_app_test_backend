package main

import (
	"backend/database"
	"backend/services"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()
	database.Init()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
        AllowOrigins: []string{"http://localhost:3000"},
        AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
    }))

	// POST
	e.POST("/create/user", services.NewUser)
	e.POST("/create/post", services.NewPost)

	// GET
	e.GET("/get/users", services.GetUsers)
	e.GET("/get/posts", services.GetPosts)

	// PUT
	e.PUT("update/username/:id", services.UpdateUser)
	e.PUT("update/post/:postId", services.UpdatePost)

	// DELETE
	e.DELETE("delete/user/:id", services.DeleteUser)
	e.DELETE("delete/post/:postId", services.DeletePost)

	e.Logger.Fatal(e.Start(":8080"))
}