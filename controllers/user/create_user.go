package usercontroller

import (
	dto "Api/dto"
	"Api/entities"
	"errors"
	"gorm.io/gorm"
)

func CreateUser(user dto.SignUpDto, db *gorm.DB) (entities.User, error) {
	createdUser, err := dto.ParseSignUpDto(user)

	if err != nil {
		return createdUser, errors.New("formato de data inválido, por favor envie a data no formato '1999-12-31'")
	}

	result := db.Create(&createdUser)

	return createdUser, result.Error
}
