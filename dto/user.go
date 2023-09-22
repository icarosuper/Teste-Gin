package dto

import (
	"Api/utils"
	"time"
)

type BirthdateFromReqDTO struct {
	day   int
	month int
	year  int
}

type GetUserFromReqDTO struct {
	Name      string
	Email     string
	Birthdate BirthdateFromReqDTO
}

type CreateUserDTO struct {
	Name      string
	Email     string
	Birthdate time.Time
	role      string
}

type ReturnUserDTO struct {
	Name      string
	Email     string
	Birthdate time.Time
}

func ParseReqUser(user GetUserFromReqDTO) CreateUserDTO {
	return CreateUserDTO{
		Name:      user.Name,
		Email:     user.Email,
		Birthdate: utils.GetDate(user.Birthdate.day, time.Month(user.Birthdate.month), user.Birthdate.year),
		role:      "user",
	}
}
