package usercontroller

import (
	"Api/entities"
	"gorm.io/gorm"
)

func GetUserByEmail(email string, db *gorm.DB) (entities.User, error) {
	var user entities.User

	result := db.Where("email = ?", email).First(&user)

	return user, result.Error
}
