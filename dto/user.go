package dto

import (
	"Api/utils"
	"time"
)

type Birthdate struct {
	day   int
	month int
	year  int
}

type GetUserFromReqDTO struct {
	Name  string
	Email string
	Birthdate
}

type CreateUserDTO struct {
	Name      string
	Email     string
	Birthdate time.Time
	role      string
}

type ReturnUserDTO struct {
	Name      string
	Birthdate time.Time
}

func ParseReqUser(user GetUserFromReqDTO) CreateUserDTO {
	return CreateUserDTO{
		Name:      user.Name,
		Email:     user.Email,
		Birthdate: utils.GetDate(user.day, time.Month(user.month), user.year),
		role:      "user",
	}
}
