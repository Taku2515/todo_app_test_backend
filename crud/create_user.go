package crud

import (
	"backend/database"
)

func CreateUserDB(user database.User) database.User {

	database.DB.Create(&user)
	return user

}