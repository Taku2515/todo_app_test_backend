package crud

import (
	"backend/database"
)

func FetchPosts(posts []database.Post) []database.Post {

	database.DB.Find(&posts)
	return posts

}