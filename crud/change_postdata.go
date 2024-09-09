package crud

import (
	"backend/database"
)

func ChangePost(post database.Post, newTask string, importance uint, deadline string) database.Post {

	database.DB.Model(&post).Update("task", newTask)
	database.DB.Model(&post).Update("importance", importance)
	database.DB.Model(&post).Update("deadline", deadline)

	return post

}