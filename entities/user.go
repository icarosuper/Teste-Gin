package entities

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	Id        uuid.UUID `gorm:"primaryKey; ->; type: uuid NOT NULL UNIQUE DEFAULT gen_random_uuid() PRIMARY KEY"`
	Name      string    `gorm:"<-; type: VARCHAR (50) NOT NULL"`
	Email     string    `gorm:"<-; type: VARCHAR (50) NOT NULL"`
	Birthdate time.Time
	Role      string
	CreatedAt time.Time `gorm:"index; <-create"`
	UpdatedAt time.Time
}
