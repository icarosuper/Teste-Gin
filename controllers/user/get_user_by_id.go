package usercontroller

import (
	"Api/entities"
	"gorm.io/gorm"
)

func GetUserById(id int, db *gorm.DB) (entities.User, error) {
	var user entities.User

	result := db.Find(&user, id)

	return user, result.Error
}
