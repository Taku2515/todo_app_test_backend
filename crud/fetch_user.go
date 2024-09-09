package crud

import (
	"backend/database"
)

func FetchUser(userID uint) (database.User, error) {

	var user database.User

	if err := database.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil

}