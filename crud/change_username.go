package crud

import (
	"backend/database"
)

func ChangeUsername(user database.User, newName string) database.User {

	database.DB.Model(&user).Update("name", newName)
	return user

}