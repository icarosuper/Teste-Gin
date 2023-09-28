package usercontroller

import (
	"Api/entities"
	"gorm.io/gorm"
)

func GetUsers(db *gorm.DB) ([]entities.User, error) {
	var users []entities.User

	result := db.Find(&users)

	return users, result.Error
}
