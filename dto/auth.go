package dto

import (
	"Api/entities"
	"time"
)

type SignUpDto struct {
	Email     string
	Name      string
	Birthdate string
	Password  string
}

type SignInDto struct {
	Email    string
	Password string
}

func ParseSignUpDto(user SignUpDto) (entities.User, error) {
	birthdate, err := time.Parse("2006-01-02", user.Birthdate)

	return entities.User{
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		Birthdate: birthdate,
		Role:      "user",
	}, err
}
