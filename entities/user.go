package entities

import (
	"time"
)

type User struct {
	Uid       string
	Name      string
	Email     string
	Birthdate time.Time
	role      string
}
