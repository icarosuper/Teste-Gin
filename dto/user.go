package dto

import (
	"Api/entities"
	"time"
)

type UserDTO struct {
	Name      string
	Email     string
	Birthdate string
}

func ParseUserDTO(user UserDTO) (entities.User, error) {
	birthdate, err := time.Parse("2006-01-02", user.Birthdate)

	return entities.User{
		Name:      user.Name,
		Email:     user.Email,
		Birthdate: birthdate,
		Role:      "user",
	}, err
}
