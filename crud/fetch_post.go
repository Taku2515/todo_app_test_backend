package crud

import (
	"backend/database"
)

func FetchPost(postID uint) (database.Post, error) {

	var post database.Post

	if err := database.DB.Where("id = ?", postID).First(&post).Error; err != nil {
		return post, err
	}
	return post, nil

}