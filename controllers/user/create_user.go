package usercontroller

import (
	dto "Api/dto"
	"errors"
	"gorm.io/gorm"
)

func CreateUser(user dto.UserDTO, db *gorm.DB) error {
	createdUser, err := dto.ParseUserDTO(user)

	if err != nil {
		return errors.New("formato de data inválido, por favor envie a data no formato '1999-12-31'")
	}

	result := db.Create(&createdUser)

	return result.Error
}
