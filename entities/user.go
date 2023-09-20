package entities

import (
	"Api/utils"
	"time"
)

type User struct {
	name     string
	birth    time.Time
	userType string
}

func NewUser(name string) User {
	return User{
		name,
		utils.GetDate(21, time.Month(6), 2000),
		"admin",
	}
}
