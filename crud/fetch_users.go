package crud

import (
	"backend/database"
)

func FetchUsers(users []database.User) []database.User {

	database.DB.Find(&users)
	return users

}