package crud

import (
	"backend/database"
)

func RemoveUser(userID uint) database.User {

	user := database.User{}
	database.DB.Where("id = ?", userID).Delete(&user)
	return user

}