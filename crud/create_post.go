package crud

import (
	"backend/database"
)

func CreatePostDB(post database.Post) database.Post {

	database.DB.Create(&post)
	return post

}